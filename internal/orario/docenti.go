package orario

import (
	"fmt"
)

type docenti []Docente
type Docente struct {
	Nome    string `json:"nome"`
	Cognome string `json:"cognome"`
	Data    string `json:"encrypted,omitempty"`
}

var (
	doc docenti
)

func loadDocenti() {
	doc = nil
	var doctemp docenti
	for _, att := range orario.Attivita {
		doctemp = append(doctemp, Docente{att.DocNome, att.DocCognome, ""})
	}

	for _, d := range doctemp {
		skip := false
		for _, u := range doc {
			if d == u {
				skip = true
				break
			}
		}
		if !skip {
			doc = append(doc, d)
		}
	}
}

func GetAllDocenti() *docenti {
	return &doc
}

func (doc *Docente) Validate() error {
	if doc.Nome == "" {
		return fmt.Errorf("JSON incompleto - nome mancante")
	}
	if doc.Cognome == "" {
		return fmt.Errorf("JSON incompleto - cognome mancante")
	}
	return nil
}
