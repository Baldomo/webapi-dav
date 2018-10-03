package orario

import (
	"encoding/xml"
	"github.com/Baldomo/webapi-dav/log"
	"io/ioutil"
	"strconv"
	"strings"
)

type table struct {
	Nome     string     `xml:"nom,attr"`
	Attivita []attivita `xml:"Attivita"`
}

type attivita struct {
	Num        uint   `xml:"Numero" json:"num"`
	Durata     durata `xml:"DURATA" json:"durata,omitempty"`
	MatCod     string `xml:"MAT_COD" json:"mat_cod,omitempty"`
	Materia    string `xml:"MAT_NOME" json:"materia,omitempty"`
	DocCognome string `xml:"DOC_COGN" json:"doc_cognome,omitempty"`
	DocNome    string `xml:"DOC_NOME" json:"doc_nome,omitempty"`
	Classe     classe `xml:"CLASSE" json:"classe,omitempty"`
	Aula       string `xml:"AULA" json:"aula,omitempty"`
	Giorno     string `xml:"GIORNO" json:"giorno,omitempty"`
	Inizio     inizio `xml:"O.INIZIO" json:"inizio,omitempty"`
	Sede       string `xml:"SEDE" json:"sede,omitempty"`
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
	if content == "" {
		// empty tag
		return nil
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
	var err error
	if err = decoder.DecodeElement(&content, &start); err != nil {
		return err
	}

	if content == "" {
		// empty tag
		return nil
	}

	var meridian = "AM"
	if h, err := strconv.Atoi(strings.Split(content, "h")[1]); err == nil {
		if h >= 12 {
			meridian = "PM"
		}
	} else {
		return err
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
		if (strings.ToLower(att.DocCognome) == strings.ToLower(doc.Cognome) || doc.Cognome == "") &&
			(strings.ToLower(att.DocNome) == strings.ToLower(doc.Nome) || doc.Nome == "") {
			a = append(a, att)
		}
	}
	return &a
}
