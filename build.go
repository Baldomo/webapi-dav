//+build build

package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	fast   bool
	goos   string
	outDir string
	run    bool

	help bool
)

type cmd func()

// These are the commands made available through the CLI
var commands = map[string]cmd{
	"build":  build,
	"clean":  clean,
	"deploy": deploy,
	"docker": docker,
	"test":   test,
}

var descriptions = map[string]string{
	"build":  "build and package artifacts",
	"clean":  "cleanup build output and playground/",
	"deploy": "create a standard working environment to test the program locally",
	"docker": "run webapi-dav in a Docker container",
	"test":   "run tests for all packages",
}

func main() {
	flag.BoolVar(&fast, "fast", false, "skip archiving builds in .tar.gz files and checksum generation")
	flag.StringVar(&goos, "os", "windows:linux", "systems to build for (separated by column, e.g. `windows:linux:mac`)")
	flag.StringVar(&outDir, "out", "build", "specifies build output `directory`")
	flag.BoolVar(&run, "run", false, "run webapi after deployment")
	flag.BoolVar(&help, "help", false, "prints usage")
	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage: go run build.go [OPTIONS] <COMMAND>\n")
		fmt.Printf("\nCommands:\n")
		for param, desc := range descriptions {
			fmt.Printf("  %s\n    \t%s\n", param, desc)
		}
		fmt.Printf("\nFlags:\n")
		flag.PrintDefaults()
	}

	command := flag.Arg(0)
	if command == "" {
		flag.Usage()
		return
	}

	if invoked, ok := commands[command]; ok {
		invoked()
	} else {
		fmt.Printf("Unknown command: %s\n", command)
		flag.Usage()
	}
}

func must(err error) {
	if err != nil {
		fmt.Printf("err: %s\n", err)
		panic(err)
	}
}

// Builds binaries for the various OS's
func build() {
	vars := map[string]string{
		"windows": "windows",
		"linux":   "linux",
		"mac":     "darwin",
	}

	osList := strings.Split(goos, ":")
	for _, osValue := range osList {
		if v, ok := vars[osValue]; ok {
			buildArtifact(v)
			if !fast {
				packageArtifact(v)
				writeChecksum(v)
			}
		} else {
			log.Printf("Invalid OS supplied: %s, skipping...\n", osValue)
		}
	}
}

func buildArtifact(system string) {
	os.Setenv("GOOS", system)
	os.Setenv("GOARCH", "amd64")
	os.Setenv("CGO_ENABLED", "0")

	fileExt := ""
	ldflags := "-s -w"
	if system == "windows" {
		fileExt = ".exe"
		ldflags += " -H windowsgui"
	}

	outFile := path.Join(outDir, system, "webapi")
	outFile += fileExt

	args := []string{
		"build",
		"-o", outFile,
		fmt.Sprint("-ldflags=", ldflags),
		"./cmd/webapi",
	}

	buildCmd := exec.Command("go", args...)
	log.Printf("Building package for %s %v\n", system, buildCmd.Args)
	buildCmd.Stderr = os.Stderr
	buildCmd.Stdout = os.Stdout
	err := buildCmd.Run()
	must(err)
}

func packageArtifact(system string) {
	binFile := path.Join(outDir, system, "webapi")
	if system == "windows" {
		binFile += ".exe"
	}

	// outDir/webapi-linux.tar.gz
	archiveName := fmt.Sprintf("%s-%s.tar.gz", "webapi", system)
	archivePath := path.Join(outDir, archiveName)
	file, err := os.OpenFile(archivePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	must(err)

	gzWriter, err := gzip.NewWriterLevel(file, gzip.BestCompression)
	defer gzWriter.Close()
	must(err)

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	// Bools define wether to keep the original file path (true) or just the file name (false
	files := map[string]bool{
		binFile:             false,
		"config.toml":       true,
		"orario.xml":        true,
		"docs/index.html":   true,
		"docs/openapi.yaml": true,
	}

	log.Printf("tar-zipping artifact %s for %s\n", archiveName, system)

	for filename, keepPath := range files {
		file, err := os.Open(filename)
		must(err)
		info, err := file.Stat()
		must(err)

		if !keepPath {
			filename = path.Base(filename)
		}

		hdr := &tar.Header{
			Name:    filename,
			Mode:    int64(info.Mode()),
			Size:    info.Size(),
			ModTime: info.ModTime(),
		}

		err = tarWriter.WriteHeader(hdr)
		must(err)

		_, err = io.Copy(tarWriter, file)
		must(err)
	}
}

func writeChecksum(system string) {
	checksumPath := path.Join(outDir, "checksums.txt")
	checksumFile, err := os.OpenFile(checksumPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer checksumFile.Close()
	must(err)

	archiveName := fmt.Sprintf("%s-%s.tar.gz", "webapi", system)
	archivePath := path.Join(outDir, archiveName)
	archiveFile, err := os.Open(archivePath)
	defer archiveFile.Close()
	must(err)

	hasher := sha256.New()
	_, err = io.Copy(hasher, archiveFile)
	must(err)

	log.Printf("Writing checksum for file %s", archiveName)
	_, err = checksumFile.WriteString(fmt.Sprintf("%x  %s\n", hasher.Sum(nil), archiveName))
	must(err)
}

// Removes built artifacts and deployment directory (playground)
func clean() {
	paths := []string{
		outDir,
		"playground",
	}

	for _, p := range paths {
		err := os.RemoveAll(p)
		must(err)
	}
}

// Creates a dummy deploy in a local directory, with dummy files and stuff
func deploy() {
	// drwxr-xr-x
	err := os.Mkdir("playground", os.ModePerm)
	must(err)
	dummyFiles(100,
		"playground/comunicati-genitori",
		"playground/comunicati-studenti",
		"playground/comunicati-docenti",
	)

	smartCopy("docs/", "playground/docs/")
	smartCopy("config.toml", "playground/config.toml")
	smartCopy("orario.xml", "playground/orario.xml")

	fileExt := ""
	if runtime.GOOS == "windows" {
		fileExt = ".exe"
	}
	binFile := path.Join(outDir, runtime.GOOS, "webapi")
	binFile += fileExt
	smartCopy(binFile, "playground/webapi"+fileExt)

	if run {
		apiCmd := exec.Command("./webapi" + fileExt)
		apiCmd.Dir = "playground"
		log.Println("Starting webapi")
		log.Println("Use Ctrl+C (SIGINT) to exit...")
		err := apiCmd.Run()
		must(err)
	}
}

func dummyFiles(num int, dirs ...string) {
	if len(dirs) == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	for _, dir := range dirs {
		err := os.Mkdir(dir, os.ModePerm)
		must(err)

		for i := 0; i <= num; i++ {
			filename := fmt.Sprintf("%s/%d.pdf", dir, i)
			file, err := os.Create(filename)
			must(err)
			err = file.Truncate(randInt64(4e5, 4e6))
			defer file.Close()
			must(err)
		}
	}
}

func randInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func smartCopy(src, dest string) {
	info, err := os.Lstat(src)
	must(err)

	if info.IsDir() {
		dirCopy(src, dest, info)
		return
	}
	fileCopy(src, dest, info)
}

func fileCopy(src, dest string, info os.FileInfo) {
	err := os.MkdirAll(filepath.Dir(dest), os.ModePerm)
	must(err)

	f, err := os.Create(dest)
	defer f.Close()
	must(err)

	err = os.Chmod(f.Name(), info.Mode())
	must(err)

	s, err := os.Open(src)
	defer s.Close()
	must(err)

	_, err = io.Copy(f, s)
	must(err)
}

func dirCopy(srcdir, destdir string, info os.FileInfo) {
	originalMode := info.Mode()

	// Make dest dir with 0755 so that everything is writable
	err := os.MkdirAll(destdir, os.FileMode(0755))
	must(err)
	// Recover dir mode with original one
	defer os.Chmod(destdir, originalMode)

	contents, err := ioutil.ReadDir(srcdir)
	must(err)

	for _, content := range contents {
		contentSrc, contentDest := filepath.Join(srcdir, content.Name()), filepath.Join(destdir, content.Name())
		smartCopy(contentSrc, contentDest)
	}
}

// Deploys a Docker container with webapi-dav
func docker() {
	argsImages := []string{
		"images", "-q",
		"-f", "reference=webapi-dav",
	}
	imagesRaw, err := exec.Command("docker", argsImages...).Output()
	must(err)
	if len(imagesRaw) == 0 {
		log.Println("Docker image not found! Building image webapi-dav")
		argsBuild := []string{
			"image", "build",
			"-t", "webapi-dav",
			".",
		}

		buildCmd := exec.Command("docker", argsBuild...)
		buildCmd.Stdout = os.Stdout
		must(buildCmd.Run())
	}

	argsPs := []string{
		"ps", "-q", "-a",
		"-f", "ancestor=webapi-dav",
	}
	containersRaw, err := exec.Command("docker", argsPs...).Output()
	must(err)
	if len(containersRaw) != 0 {
		// There are running webapi-dav instances
		log.Println("Stopping running instances")
		latest := strings.Split(string(containersRaw), "\n")[0]

		argsStop := []string{
			"stop",
			latest,
		}
		must(exec.Command("docker", argsStop...).Run())

		log.Println("Removing existing containers")
		argsRm := []string{
			"container", "rm", "webapi-dav",
		}
		must(exec.Command("docker", argsRm...).Run())
	}

	comDoc, _ := filepath.Abs("./playground/comunicati-docenti")
	comGen, _ := filepath.Abs("./playground/comunicati-genitori")
	comStud, _ := filepath.Abs("./playground/comunicati-studenti")

	_ = os.MkdirAll(comDoc, os.ModeDir|os.ModePerm)
	_ = os.MkdirAll(comGen, os.ModeDir|os.ModePerm)
	_ = os.MkdirAll(comStud, os.ModeDir|os.ModePerm)

	log.Println("Starting container")
	argsRun := []string{
		"run",
		"-it", "-d",
		"--name", "webapi-dav",
		"-p", "8080:8080",
		"-v", comDoc + ":/comunicati-docenti",
		"-v", comGen + ":/comunicati-genitori",
		"-v", comStud + ":/comunicati-studenti",
		"webapi-dav",
	}
	runCmd := exec.Command("docker", argsRun...)
	runCmd.Stdout = os.Stdout
	must(runCmd.Run())
}

// Runs tests
func test() {
	args := []string{
		"test",
		"-v",
		"-benchmem",
		"./...",
	}

	testCmd := exec.Command("go", args...)
	testCmd.Stdout = os.Stdout
	must(testCmd.Run())
}
