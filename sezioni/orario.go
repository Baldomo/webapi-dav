package sezioni

import (
	"encoding/xml"
	"io/ioutil"
	"leonardobaldin/webapi-dav/log"
	"regexp"
	"strings"
)

type table struct {
	Nome     string     `xml:"nom,attr"`
	Attivita []attivita `xml:"Attivita"`
}

type attivita struct {
	Num        uint   `xml:"Numero" json:"num"`
	Durata     durata `xml:"DURATA" json:"durata"`
	MatCod     string `xml:"MAT_COD" json:"mat_cod"`
	Materia    string `xml:"MAT_NOME" json:"materia"`
	DocCognome string `xml:"DOC_COGN" json:"doc_cognome"`
	DocNome    string `xml:"DOC_NOME" json:"doc_nome"`
	Classe     classe `xml:"CLASSE" json:"classe,omitempty"`
	Aula       string `xml:"AULA" json:"aula"`
	Giorno     string `xml:"GIORNO" json:"giorno"`
	Inizio     inizio `xml:"O.INIZIO" json:"inizio"`
	Sede       string `xml:"SEDE" json:"sede"`
}

type docenti []docente
type docente struct {
	Nome    string `xml:"TABLE>Attivita>DOC_NOME"`
	Cognome string `xml:"TABLE>Attivita>DOC_COGN"`
}

type durata string
type classe string
type inizio string

var (
	orario *table
	doc    docenti

	reClasse = regexp.MustCompile("[1-5][a-zA-Z]")
)

func (c *classe) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := decoder.DecodeElement(&content, &start); err != nil {
		return err
	}
	content = reClasse.FindString(content)
	*c = classe(content)
	return nil
}

func (c classe) String() string {
	return string(c)
}

func (d *durata) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := decoder.DecodeElement(&content, &start); err != nil {
		return err
	}
	if !strings.HasSuffix(content, "m") {
		content += "m"
	}
	*d = durata(content)
	return nil
}

func (d durata) String() string {
	return string(d)
}

func (i *inizio) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := decoder.DecodeElement(&content, &start); err != nil {
		return err
	}
	content = strings.Replace(content, "h", ":", -1) + "AM"
	*i = inizio(content)
	return nil
}

func (i inizio) String() string {
	return string(i)
}

func LoadOrario(path string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Log.Error(err.Error())
		return
	}
	if err := xml.Unmarshal(raw, &orario); err != nil {
		log.Log.Error(err.Error())
		return
	}

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

func GetOrario() *table {
	return orario
}

func GetByGiorno(giorno string) *[]attivita {
	var a []attivita
	for _, att := range orario.Attivita {
		if strings.ToLower(att.Giorno) == strings.ToLower(giorno) {
			a = append(a, att)
		}
	}
	return &a
}

func GetByClasse(classe string) *[]attivita {
	var a []attivita
	for _, att := range orario.Attivita {
		if strings.ToLower(att.Classe.String()) == strings.ToLower(classe) {
			a = append(a, att)
		}
	}
	return &a
}

func GetByDocNome(nome string) *[]attivita {
	var a []attivita
	for _, att := range orario.Attivita {
		if strings.ToLower(att.DocNome) == strings.ToLower(nome) {
			a = append(a, att)
		}
	}
	return &a
}

func GetByDocCogn(cogn string) *[]attivita {
	var a []attivita
	for _, att := range orario.Attivita {
		if strings.ToLower(att.DocCognome) == strings.ToLower(cogn) {
			a = append(a, att)
		}
	}
	return &a
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
