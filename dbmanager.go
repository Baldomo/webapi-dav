package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const(
	agendaTable = "npjmx_jevents_vevdetail"
	timeFormat = "0000-00-00 00:00:00"
)

type Event struct {
	Date    []time.Time `json:"date" xml:"Date"`
	Title   string		`json:"titolo" xml:"Titolo"`
	Content string		`json:"contenuto" xml:"Contenuto"`
}

func GetAgenda() error {
	db, err := sql.Open("mysql", "leonardo:leo@/agenda")
	if err != nil {
		return err
	}
	db.Close()

	return nil
}

func (e *Event) FillMissingInfo() (*Event, error) {
	if len(e.Date) == 0 {

	}

	return &Event{}, nil
}