package comunicati

import (
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/Baldomo/webapi-dav/internal/config"
	"github.com/Baldomo/webapi-dav/internal/log"
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

	UrlPrefix = "http://www.liceodavinci.tv/sitoLiceo/images/comunicati/"
)

func (c *Comunicato) Equals(other *Comunicato) bool {
	if c.Nome == other.Nome && c.Data.Round(time.Second).Equal(other.Data.Round(time.Second)) && c.Tipo == other.Tipo && c.URL == other.URL {
		return true
	}
	return false
}

func NewComunicato(nome string, data time.Time, tipo string) *Comunicato {
	com := new(Comunicato)
	com.Nome = nome
	com.Data = data
	if strings.ContainsAny(tipo, "/") {
		tipo = strings.Replace(tipo, "/", "", -1)
	}
	com.Tipo = tipo
	com.URL = UrlPrefix + "comunicati-" + tipo + "/" + url.PathEscape(nome)

	return com
}

func scrape(dir string, tipo string) Comunicati {
	var wg sync.WaitGroup

	absPath, _ := filepath.Abs(dir)
	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		log.Fatal(err)
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
	log.Infof("Caricamento comunicati %s completato", tipo)
	return buf
}

func LoadComunicati(tipo string) {
	switch tipo {

	case TipoGenitori:
		Genitori = scrape(config.GetConfig().Dirs.Genitori, TipoGenitori)

	case TipoDocenti:
		Docenti = scrape(config.GetConfig().Dirs.Docenti, TipoDocenti)

	case TipoStudenti:
		Studenti = scrape(config.GetConfig().Dirs.Studenti, TipoStudenti)

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
