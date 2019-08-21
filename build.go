///+build build

package main

import (
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
	out_dir string
	run     bool
	goos    string
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
	flag.StringVar(&out_dir, "out", "build", "specifies build output `directory`")
	flag.BoolVar(&run, "run", false, "run webapi after deployment")
	flag.StringVar(&goos, "os", "windows:linux", "systems to build for (separated by column, e.g. `windows:linux:mac`)")
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
			err := buildArtifact(v)
			if err != nil {
				return err
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

	file_ext := ""
	ldflags := "-s -w"
	if system == "windows" {
		file_ext = ".exe"
		ldflags += " -H windowsgui"
	}

	out_file := path.Join(out_dir, system, "webapi")
	out_file += file_ext

	args := []string{
		"build",
		"-o", out_file,
		"-ldflags", ldflags,
		"./cmd/webapi",
	}

	buildCmd := exec.Command("go", args...)
	log.Printf("Building package for %s %v\n", system, buildCmd.Args)
	return buildCmd.Run()
}

// Removes built artifacts and deployment directory (playground)
func clean() error {
	paths := []string{
		out_dir,
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

	file_ext := ""
	if runtime.GOOS == "windows" {
		file_ext = ".exe"
	}
	bin_file := path.Join(out_dir, runtime.GOOS, "webapi")
	bin_file += file_ext
	err = smartCopy(bin_file, "playground/webapi" + file_ext)
	if err != nil {
		return err
	}

	if run {
		apiCmd := exec.Command("./webapi" + file_ext)
		apiCmd.Dir = "playground"
		log.Printf("Starting webapi\nUse Ctrl+C (SIGINT) to exit...")
		apiCmd.Start()
		apiCmd.Wait()
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
	return testCmd.Run()
}
