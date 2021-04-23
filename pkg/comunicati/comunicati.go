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

	"github.com/Baldomo/webapi-dav/pkg/config"
	"github.com/Baldomo/webapi-dav/pkg/log"
)

// Rappresentazione di un comunicato come restituito dalle richieste REST
type Comunicato struct {
	// Nome/titolo del comunicato (praticamente il nome del file)
	Nome string `json:"nome"`

	// Data di emissione del comunicato
	Data time.Time `json:"data"`

	// Tipo del comunicato (docenti/genitori/studenti)
	Tipo string `json:"tipo"`

	// URL statico del comunicato sul server
	URL string `json:"url"`
}

type Comunicati []*Comunicato

// Variabili interne di default per tenere in memoria i comunicati dei vari tipi
var (
	Genitori Comunicati
	Studenti Comunicati
	Docenti  Comunicati
)

// Costanti di default per i tipi dei comunicati e l'URL di base della cartella
// contenente i file PDF
const (
	TipoGenitori = "genitori"
	TipoStudenti = "studenti"
	TipoDocenti  = "docenti"

	PathPrefix = "/sitoLiceo/images/comunicati/"
)

// Compara due comunicati con arrotondamento della data etc
func (c *Comunicato) Equals(other *Comunicato) bool {
	if c.Nome == other.Nome &&
		c.Data.Round(time.Second).Equal(other.Data.Round(time.Second)) &&
		c.Tipo == other.Tipo &&
		c.URL == other.URL {
		return true
	}
	return false
}

// Costruttore per Comunicato
func NewComunicato(nome string, data time.Time, tipo string) *Comunicato {
	com := new(Comunicato)
	com.Nome = nome
	com.Data = data
	if strings.ContainsAny(tipo, "/") {
		tipo = strings.Replace(tipo, "/", "", -1)
	}
	com.Tipo = tipo
    com.URL = "https://" + config.GetConfig().General.FQDN + PathPrefix + "comunicati-" + tipo + "/" + url.PathEscape(nome)

	return com
}

// Cammina le cartelle dei comunicati e memeorizza le informazioni di modello Comunicato
// per ogni file
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

// Carica in memoria i vari tipi di comunicati
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

// Dato un tipo di comunicato (docenti/genitori/studenti), restituisce la slice
// di Comunicato memorizzata
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

// Dato un tipo di comunicato, restituisce il numero di comunicati presenti di quel tipo
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
