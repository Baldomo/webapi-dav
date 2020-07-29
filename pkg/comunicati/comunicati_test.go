package comunicati

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	NUM_FILES_TEST = 2
	FILE_SIZE      = 4 * 1024
)

var NumFilesBenchmark = []int{
	120, 250, 500, 1000,
}

var dirs = []string{
	os.TempDir() + "/comunicati-genitori",
	os.TempDir() + "/comunicati-docenti",
	os.TempDir() + "/comunicati-studenti",
}

func TestLoadComunicati(t *testing.T) {
	// TODO: refactor
	t.SkipNow()

	var exp, got Comunicati

	rand.Seed(time.Now().UnixNano())
	data := make([]byte, FILE_SIZE)
	for _, d := range dirs {
		err := os.Mkdir(d, os.ModePerm)
		if os.IsExist(err) {
			os.RemoveAll(d)
			os.Mkdir(d, os.ModePerm)
		}

		for ind := 0; ind <= NUM_FILES_TEST; ind++ {
			filename := fmt.Sprintf("%d.txt", ind)
			file, err := os.OpenFile(d+"/"+filename, os.O_WRONLY|os.O_CREATE, os.ModePerm)
			if err != nil {
				t.Fatalf("error writing files: \n %s", err.Error())
			}
			rand.Read(data)
			file.Write(data)
			exp = append(exp, NewComunicato(
				filename,
				time.Now(),
				strings.Split(d, "-")[1],
			))
		}

		got = append(got, scrape(d, strings.Split(d, "-")[1])...)
	}

	if !areEqual(exp, got, t) {
		t.Fail()
	}
}

func areEqual(c1, c2 Comunicati, t *testing.T) bool {
	if len(c1) != len(c2) {
		t.Logf("Different lengths:\n%+v\n%+v", len(c1), len(c2))
		return false
	}

	for i := range c1 {
		if !c1[i].Equals(c2[i]) {
			t.Logf("Different comunicati:\n%+v\n%+v", c1[i], c2[i])
			return false
		}
	}

	return true
}

func BenchmarkLoadComunicati(b *testing.B) {
	// TODO
}
