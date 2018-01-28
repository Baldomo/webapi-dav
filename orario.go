package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Table struct {
	Nome     string
	Attivita []Attivita
}

type xmlTable struct {
	Nome     string        `xml:"nom,attr"`
	Attivita []xmlAttivita `xml:"Attivita"`
}

type Attivita struct {
	Num        uint
	Durata     time.Duration
	MatCod     string
	Materia    string
	DocCognome string
	DocNome    string
	Classe     string
	Aula       string
	Giorno     string
	Inizio     time.Time
	Sede       string
}

type xmlAttivita struct {
	Num        uint   `xml:"Numero"`
	Durata     string `xml:"DURATA"`
	MatCod     string `xml:"MAT_COD"`
	Materia    string `xml:"MAT_NOME"`
	DocCognome string `xml:"DOC_COGN"`
	DocNome    string `xml:"DOC_NOME"`
	Classe     string `xml:"CLASSE"`
	Aula       string `xml:"AULA"`
	Giorno     string `xml:"GIORNO"`
	Inizio     string `xml:"O.INIZIO"`
	Sede       string `xml:"SEDE"`
}

var (
	Orario    *Table
	xmlOrario xmlTable
)

func (a xmlAttivita) Attivita() Attivita {
	if !strings.HasSuffix(a.Durata, "m") {
		a.Durata += "m"
	}
	durata, errDur := time.ParseDuration(a.Durata)
	if errDur != nil {
		return Attivita{}
	}

	inizio, errIn := time.Parse("08h00", a.Inizio)
	if errIn != nil {
		return Attivita{}
	}

	return Attivita{
		a.Num,
		durata,
		a.MatCod,
		a.Materia,
		a.DocCognome,
		a.DocNome,
		a.Classe,
		a.Aula,
		a.Giorno,
		inizio,
		a.Sede,
	}
}

func (o xmlTable) Table() (t *Table) {
	t = &Table{}
	t.Nome = o.Nome
	t.Attivita = make([]Attivita, len(o.Attivita))
	for i, attivita := range o.Attivita {
		t.Attivita[i] = attivita.Attivita()
	}
	return
}

func LoadOrario(path string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		Log.Error(err.Error())
		return
	}

	if err := xml.Unmarshal(raw, &xmlOrario); err != nil {
		Log.Error(err.Error())
		return
	}
	fmt.Println(xmlOrario)

	Orario = xmlOrario.Table()
}
