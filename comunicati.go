package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

type Comunicato struct {
	Nome string    `json:"nome"`
	Data time.Time `json:"data"`
	Tipo string    `json:"tipo"`
	URL  string    `json:"url"`
}

type Comunicati []*Comunicato

var (
	Genitori Comunicati
	Studenti Comunicati
	Docenti  Comunicati
)

const (
	TipoGenitori = "genitori"
	TipoStudenti = "studenti"
	TipoDocenti  = "docenti"

	urlPrefix = "http://liceodavinci.tv/sitoLiceo/images/comunicati/"
)

func NewComunicato(nome string, data time.Time, tipo string) *Comunicato {
	com := new(Comunicato)
	com.Nome = nome
	com.Data = data
	if strings.ContainsAny(tipo, "/") {
		strings.Replace(tipo, "/", "", -1)
	}
	com.Tipo = tipo
	com.URL = urlPrefix + "comunicati-" + tipo + "/" + com.Nome

	return com
}

func fetchComunicati(dir string, tipo string) Comunicati {
	var wg sync.WaitGroup

	absPath, _ := filepath.Abs(dir)
	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		Log.Fatal(err)
	}

	buf := make(Comunicati, len(files))
	sem := make(chan int, 10)
	for i, file := range files {
		wg.Add(1)
		go func(i int, file os.FileInfo) {
			sem <- 1
			buf[i] = NewComunicato(files[i].Name(), files[i].ModTime(), tipo)
			wg.Done()
			runtime.Gosched()
			<-sem
		}(i, file)
	}

	wg.Wait()
	sort.Slice(buf, func(i, j int) bool {
		return buf[i].Data.After(buf[j].Data)
	})
	Log.Infof("Caricamento comunicati %s completato", tipo)
	return buf
}

func LoadComunicati(tipo string) {
	switch tipo {

	case TipoGenitori:
		Genitori = fetchComunicati(GetConfig().Dirs.Genitori, TipoGenitori)

	case TipoDocenti:
		Docenti = fetchComunicati(GetConfig().Dirs.Docenti, TipoDocenti)

	case TipoStudenti:
		Studenti = fetchComunicati(GetConfig().Dirs.Studenti, TipoStudenti)

	default:
		return

	}
}

func GetByName(tipo string) Comunicati {
	switch tipo {

	case "genitori":
		return Genitori

	case "studenti":
		return Studenti

	case "docenti":
		return Docenti

	default:
		return nil
	}
}

func GetLenByName(tipo string) int {
	switch tipo {

	case "genitori":
		return len(Genitori)

	case "studenti":
		return len(Studenti)

	case "docenti":
		return len(Docenti)

	default:
		return 0
	}
}
