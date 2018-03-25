package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

const (
	agendaTable = "npjmx_jevents_vevdetail"
	// timeFormat  = "0000-00-00 00:00:00"
	// time.Unix(epoch, 0).Format(time.RFC822Z)

	dataSource = "leonardo:leonardo@/agenda"

	titleField   = "summary"
	contentField = "description"
	inizioField  = "dtstart"
	fineField    = "dtend"
)

var (
	db *sqlx.DB
)

type Date struct {
	Inizio time.Time
	Fine   time.Time
}

type Event struct {
	Title   string `json:"titolo" xml:"Titolo"`
	Content string `json:"contenuto" xml:"Contenuto"`
	Date    `json:"date" xml:"Date"`
}

func GetAgenda() error {
	var err error
	db, err = sqlx.Open("mysql", dataSource)
	if err != nil {
		db.Close()
		return err
	}
	db.Close()

	return nil
}

func (e *Event) FillEmptyFields() error {
	query := buildQuery(e)
	fmt.Println(query)

	// errore se risultati query > 1
	return nil
}

func buildQuery(e *Event) string {
	var emptyFields []string
	var keyFields []string
	if e.Title == "" {
		emptyFields = append(emptyFields, titleField)
	} else {
		keyFields = append(keyFields, titleField+"=:"+titleField)
	}
	if e.Content == "" {
		emptyFields = append(emptyFields, contentField)
	} else {
		keyFields = append(keyFields, contentField+"=:"+contentField)
	}
	if e.Date.Inizio.IsZero() {
		emptyFields = append(emptyFields, inizioField)
	} else {
		keyFields = append(keyFields, inizioField+"=:"+inizioField)
	}
	if e.Date.Fine.IsZero() {
		emptyFields = append(emptyFields, fineField)
	} else {
		keyFields = append(keyFields, fineField+"=:"+fineField)
	}

	if len(emptyFields) == 0 {
		return ""
	}

	return "select " + strings.Join(emptyFields, ",") +
		" from " + agendaTable +
		" where " + strings.Join(keyFields, " and ")
}
