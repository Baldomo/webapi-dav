package orario

import (
	"encoding/xml"
	"io/ioutil"
	"leonardobaldin/webapi-dav/log"
	"strconv"
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

type durata string
type inizio string

var (
	orario *table
)

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

	var meridian = "AM"
	if h, errS := strconv.Atoi(content[:2]); errS != nil {
		return errS
	} else {
		if h >= 12 {
			meridian = "PM"
		}
	}
	content = strings.Replace(content, "h", ":", -1) + meridian
	*i = inizio(content)
	return nil
}

func (i inizio) String() string {
	return string(i)
}

func LoadOrario(path string) {
	orario = nil
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Log.Error(err.Error())
		return
	}
	if err := xml.Unmarshal(raw, &orario); err != nil {
		log.Log.Error(err.Error())
		return
	}

	loadDocenti()
	loadClassi()
}

func GetOrario() *table {
	return orario
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

func GetByDoc(doc Docente) *[]attivita {
	var a []attivita
	for _, att := range orario.Attivita {
		if (strings.ToLower(att.DocCognome) == strings.ToLower(doc.Cognome)) && (strings.ToLower(att.DocNome) == strings.ToLower(doc.Nome)) {
			a = append(a, att)
		}
	}
	return &a
}