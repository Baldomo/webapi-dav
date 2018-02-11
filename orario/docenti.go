package orario

import "strings"

type docenti []docente
type docente struct {
	Nome    string
	Cognome string
}

var (
	doc docenti
)

func loadDocenti() {
	doc = nil
	var doctemp docenti
	for _, att := range orario.Attivita {
		doctemp = append(doctemp, docente{att.DocNome, att.DocCognome})
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

func GetDocentiCogn(cogn string) *docenti {
	var d docenti
	for _, docente := range doc {
		if strings.ToLower(docente.Cognome) == strings.ToLower(cogn) {
			d = append(d, docente)
		}
	}
	return &d
}
