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

// Memorizza tutti i docenti a partire dalle stringhe valide contenute nella
// tabella decodificata dall'XML
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

// Restituisce un puntatore alla slice di docenti estratti dall'orario
func GetAllDocenti() *docenti {
	return &doc
}

// Convalida a livello JSON le richieste REST ricevute
func (doc *Docente) Validate() error {
	if doc.Nome == "" {
		return fmt.Errorf("JSON incompleto - nome mancante")
	}
	if doc.Cognome == "" {
		return fmt.Errorf("JSON incompleto - cognome mancante")
	}
	return nil
}
