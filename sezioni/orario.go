package sezioni

import (
	"encoding/xml"
	"io/ioutil"
	"leonardobaldin/webapi-dav/log"
	"strings"
)

type table struct {
	Nome     string     `xml:"nom,attr"`
	Attivita []attivita `xml:"Attivita"`
}

type attivita struct {
	Num        uint   `xml:"Numero"`
	Durata     durata `xml:"DURATA"`
	MatCod     string `xml:"MAT_COD"`
	Materia    string `xml:"MAT_NOME"`
	DocCognome string `xml:"DOC_COGN"`
	DocNome    string `xml:"DOC_NOME"`
	Classe     string `xml:"CLASSE"`
	Aula       string `xml:"AULA"`
	Giorno     string `xml:"GIORNO"`
	Inizio     inizio `xml:"O.INIZIO"`
	Sede       string `xml:"SEDE"`
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

func (i *inizio) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := decoder.DecodeElement(&content, &start); err != nil {
		return err
	}
	content = strings.Replace(content, "h", ":", -1) + "AM"
	*i = inizio(content)
	return nil
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
}

func GetOrario() *table {
	return orario
}

func GetByGiorno(giorno string) []*attivita {
	var a []*attivita
	for _, att := range orario.Attivita {
		if att.Giorno == giorno {
			a = append(a, &att)
		}
	}
	return a
}
