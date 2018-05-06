package orario

import (
	"encoding/xml"
	"regexp"
)

var (
	classi []classe

	reClasse = regexp.MustCompile("[1-5][a-zA-Z]")
)

type classe string

func loadClassi() {
	classi = nil
	var classitemp []classe
	for _, att := range orario.Attivita {
		classitemp = append(classitemp, att.Classe)
	}

	for _, c := range classitemp {
		skip := false
		for _, u := range classi {
			if c == u || c.String() == "" {
				skip = true
				break
			}
		}
		if !skip {
			classi = append(classi, c)
		}
	}
}

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

func GetAllClassi() *[]classe {
	return &classi
}
