///+build build

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
)

type cmd func() error

// These are the commands made available through the CLI
var commands = map[string]cmd{
	"build":  build,
	"clean":  clean,
	"deploy": deploy,
	"test":   test,
}

func main() {
	flag.BoolVar(&fast, "fast", false, "skip archiving builds in .tar.gz files and checksum generation")
	flag.StringVar(&goos, "os", "windows:linux", "systems to build for (separated by column, e.g. `windows:linux:mac`)")
	flag.StringVar(&outDir, "out", "build", "specifies build output `directory`")
	flag.BoolVar(&run, "run", false, "run webapi after deployment")
	flag.Parse()

	command := flag.Arg(0)
	if command == "" {
		flag.PrintDefaults()
		return
	}

	err := commands[command]()
	if err != nil {
		log.Fatalf("error executing command %s:\n%v\n", command, err.Error())
	}
}

// Builds binaries for the various OS's
func build() error {
	vars := map[string]string{
		"windows": "windows",
		"linux":   "linux",
		"mac":     "darwin",
	}

	osList := strings.Split(goos, ":")
	for _, osValue := range osList {
		if v, ok := vars[osValue]; ok {
			if err := buildArtifact(v); err != nil {
				return err
			}

			if !fast {
				if err := packageArtifact(v); err != nil {
					return err
				}

				if err := writeChecksum(v); err != nil {
					return err
				}
			}
		} else {
			log.Printf("Invalid OS supplied: %s, skipping...\n", osValue)
		}
	}

	return nil
}

func buildArtifact(system string) error {
	os.Setenv("GOOS", system)
	os.Setenv("GOARCH", "amd64")

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
		"-ldflags", ldflags,
		"./cmd/webapi",
	}

	buildCmd := exec.Command("go", args...)
	log.Printf("Building package for %s %v\n", system, buildCmd.Args)
	return buildCmd.Run()
}

func packageArtifact(system string) error {
	binFile := path.Join(outDir, system, "webapi")
	if system == "windows" {
		binFile += ".exe"
	}

	// outDir/webapi-linux.tar.gz
	archiveName := fmt.Sprintf("%s-%s.tar.gz", "webapi", system)
	archivePath := path.Join(outDir, archiveName)
	file, err := os.OpenFile(archivePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	gzWriter, err := gzip.NewWriterLevel(file, gzip.BestCompression)
	if err != nil {
		return err
	}
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	// Bools define wether to keep the original file path (true) or just the file name (false
	files := map[string]bool{
		binFile:               false,
		"config.toml":         true,
		"orario.xml":          true,
		"static/index.html":   true,
		"static/openapi.yaml": true,
	}

	log.Printf("tar-zipping artifact %s for %s\n", archiveName, system)

	for filename, keepPath := range files {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		info, err := file.Stat()
		if err != nil {
			return err
		}

		if !keepPath {
			filename = path.Base(filename)
		}

		hdr := &tar.Header{
			Name:    filename,
			Mode:    int64(info.Mode()),
			Size:    info.Size(),
			ModTime: info.ModTime(),
		}

		if err := tarWriter.WriteHeader(hdr); err != nil {
			return err
		}

		if _, err := io.Copy(tarWriter, file); err != nil {
			return err
		}
	}

	return nil
}

func writeChecksum(system string) error {
	checksumPath := path.Join(outDir, "checksums.txt")
	checksumFile, err := os.OpenFile(checksumPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer checksumFile.Close()

	archiveName := fmt.Sprintf("%s-%s.tar.gz", "webapi", system)
	archivePath := path.Join(outDir, archiveName)
	archiveFile, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer archiveFile.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, archiveFile); err != nil {
		return err
	}

	log.Printf("Writing checksum for file %s", archiveName)
	checksumFile.WriteString(fmt.Sprintf("%x  %s\n", hasher.Sum(nil), archiveName))

	return nil
}

// Removes built artifacts and deployment directory (playground)
func clean() error {
	paths := []string{
		outDir,
		"playground",
	}

	for _, p := range paths {
		if err := os.RemoveAll(p); err != nil {
			return err
		}
	}

	return nil
}

// Creates a dummy deploy in a local directory, with dummy files and stuff
func deploy() error {
	// drwxr-xr-x
	os.Mkdir("playground", os.ModePerm)
	dummyFiles(100,
		"playground/comunicati-genitori",
		"playground/comunicati-studenti",
		"playground/comunicati-docenti",
	)

	err := smartCopy("static/", "playground/static/")
	if err != nil {
		return err
	}

	err = smartCopy("config.toml", "playground/config.toml")
	if err != nil {
		return err
	}

	err = smartCopy("orario.xml", "playground/orario.xml")
	if err != nil {
		return err
	}

	fileExt := ""
	if runtime.GOOS == "windows" {
		fileExt = ".exe"
	}
	binFile := path.Join(outDir, runtime.GOOS, "webapi")
	binFile += fileExt
	err = smartCopy(binFile, "playground/webapi"+fileExt)
	if err != nil {
		return err
	}

	if run {
		apiCmd := exec.Command("./webapi" + fileExt)
		apiCmd.Dir = "playground"
		log.Printf("Starting webapi\nUse Ctrl+C (SIGINT) to exit...")
		apiCmd.Run()
	}

	return nil
}

func dummyFiles(num int, dirs ...string) {
	if len(dirs) == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	for _, dir := range dirs {
		os.Mkdir(dir, os.ModePerm)

		for i := 0; i <= num; i++ {
			filename := fmt.Sprintf("%s/%d.pdf", dir, i)
			file, _ := os.Create(filename)
			file.Truncate(randInt64(4e5, 4e6))
			file.Close()
		}
	}
}

func randInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func smartCopy(src, dest string) error {
	info, err := os.Lstat(src)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return dirCopy(src, dest, info)
	}
	return fileCopy(src, dest, info)
}

func fileCopy(src, dest string, info os.FileInfo) error {
	if err := os.MkdirAll(filepath.Dir(dest), os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = os.Chmod(f.Name(), info.Mode()); err != nil {
		return err
	}

	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	_, err = io.Copy(f, s)
	return err
}

func dirCopy(srcdir, destdir string, info os.FileInfo) error {
	originalMode := info.Mode()

	// Make dest dir with 0755 so that everything is writable
	if err := os.MkdirAll(destdir, os.FileMode(0755)); err != nil {
		return err
	}
	// Recover dir mode with original one
	defer os.Chmod(destdir, originalMode)

	contents, err := ioutil.ReadDir(srcdir)
	if err != nil {
		return err
	}

	for _, content := range contents {
		contentSrc, contentDest := filepath.Join(srcdir, content.Name()), filepath.Join(destdir, content.Name())
		if err := smartCopy(contentSrc, contentDest); err != nil {
			// If any error, exit immediately
			return err
		}
	}

	return nil
}

// Runs tests
func test() error {
	args := []string{
		"test",
		"-v",
		"-benchmem",
		"./...",
	}

	testCmd := exec.Command("go", args...)
	testCmd.Stdout = os.Stdout
	return testCmd.Run()
}
