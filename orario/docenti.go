package orario

import (
	"encoding/json"
	"errors"
)

type docenti []Docente
type Docente struct {
	Nome    string `json:"nome"`
	Cognome string `json:"cognome"`
}

var (
	doc docenti
)

func loadDocenti() {
	doc = nil
	var doctemp docenti
	for _, att := range orario.Attivita {
		doctemp = append(doctemp, Docente{att.DocNome, att.DocCognome})
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

func (doc *Docente) UnmarshalJSON(data []byte) error {
	if doc.Nome == "" {
		return errors.New("JSON incompleto - nome mancante")
	}
	if doc.Cognome == "" {
		return errors.New("JSON incompleto - cognome mancante")
	}
	json.Unmarshal(data, *doc)
	return nil
}
