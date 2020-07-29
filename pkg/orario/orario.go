package orario

import (
	"encoding/xml"
	"io/ioutil"
	"strings"

	"github.com/Baldomo/webapi-dav/pkg/log"
)

// Rappresentazione di una "tabella" dell'orario, che contiene una lista
// di attività
type table struct {
	Nome     string     `xml:"nom,attr"`
	Attivita []attivita `xml:"Attivita"`
}

// Rappresentazione d una singola "attività", cioè un blocco di un'ora
// dell'orario
type attivita struct {
	// Numero univoco associato a una certa attività
	Num uint `xml:"Numero" json:"num"`

	// Durata dell'attività in formato H'h'mm (ad es. "1h00")
	Durata string `xml:"DURATA" json:"durata,omitempty"`

	// Codice della materia associata all'attività (ad es. "MAT")
	MatCod string `xml:"MAT_COD" json:"mat_cod,omitempty"`

	// Nome della materia associata all'attività (ad es. "Filos\Storia")
	Materia string `xml:"MAT_NOME" json:"materia,omitempty"`

	// Cognome del docente
	DocCognome string `xml:"DOC_COGN" json:"doc_cognome,omitempty"`

	// Nome del docente
	DocNome string `xml:"DOC_NOME" json:"doc_nome,omitempty"`

	// Classe per cui si svolge l'attività (ad es. "5B")
	Classe classe `xml:"CLASSE" json:"classe,omitempty"`

	// Numero e piano dell'aula in cui si svolge l'attività (ad es. "56 3pagg")
	Aula string `xml:"AULA" json:"aula,omitempty"`

	// Nome del giorno in cui si svolge l'attività (ad es. "martedì")
	Giorno string `xml:"GIORNO" json:"giorno,omitempty"`

	// Ora di inizio in formato HH'h'mm (ad es. "08h05")
	Inizio string `xml:"O.INIZIO" json:"inizio,omitempty"`

	// Nome della sede in cui è locata l'aula (ad es. "Principale (fisica incl)")
	Sede string `xml:"SEDE" json:"sede,omitempty"`
}

var (
	orario *table
)

// Esegue parsing e decodifica di un file XML del tipo orario dato il suo percorso
// (può essere esportato da software proprietari di gestione dell'orario di istituto)
func LoadOrario(path string) {
	orario = nil
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error(err.Error())
		return
	}
	if err := xml.Unmarshal(raw, &orario); err != nil {
		log.Error(err.Error())
		return
	}

	loadDocenti()
	loadClassi()
}

// Restituisce un puntatore alla tabella orario interna
func GetOrario() *table {
	return orario
}

// Restituisce tutte le attività di una data classe (passata come, ad es. "5B",
// case insensitive)
func GetByClasse(classe string) *[]attivita {
	var a []attivita
	for _, att := range orario.Attivita {
		if strings.EqualFold(att.Classe.String(), classe) {
			a = append(a, att)
		}
	}
	return &a
}

// Restituisce tutte le attività di un dato docente
func GetByDoc(doc Docente) *[]attivita {
	var a []attivita
	for _, att := range orario.Attivita {
		if (strings.EqualFold(att.DocCognome, doc.Cognome) || doc.Cognome == "") &&
			(strings.EqualFold(att.DocNome, doc.Nome) || doc.Nome == "") {
			a = append(a, att)
		}
	}
	return &a
}
