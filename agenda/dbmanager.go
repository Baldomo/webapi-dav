package agenda

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"leonardobaldin/webapi-dav/utils"
	"strings"
	"time"
)

const (
	agendaTable = "npjmx_jevents_vevdetail"

	dataSource = "leonardo:leonardo@/agenda"

	titleField   = "summary"
	contentField = "description"
	inizioField  = "dtstart"
	fineField    = "dtend"

	baseQuery = "select " + titleField + "," + contentField + "," + inizioField + "," + fineField +
		" from " + agendaTable + " where "
)

var (
	db *sqlx.DB
)

type EventStream struct {
	After         int64    `json:"dopo"`
	Before        int64    `json:"prima"`
	TitleFilter   []string `json:"titolo_filtri"`
	ContentFilter []string `json:"contenuti_filtri"`
	events        []Event
}

type Date struct {
	Inizio time.Time `json:"inizio" xml:"inizio"`
	Fine   time.Time `json:"fine" xml:"fine"`
}

type Event struct {
	Title   string `json:"titolo" xml:"titolo"`
	Content string `json:"contenuto" xml:"contenuto"`
	Date    `json:"date" xml:"date"`
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

func NewEventStream() *EventStream {
	return &EventStream{
		After:         0,
		Before:        0,
		TitleFilter:   []string{},
		ContentFilter: []string{},
		events:        []Event{},
	}
}

func (es *EventStream) GetAfter(epoch int64) *EventStream {
	if es.Before != 0 && (epoch > es.Before || epoch == 0) {
		return es
	}
	es.After = epoch
	return es
}

func (es *EventStream) GetBefore(epoch int64) *EventStream {
	if epoch < es.After {
		return es
	}
	es.Before = epoch
	return es
}

func (es *EventStream) FilterTitle(filter []string) *EventStream {
	if len(filter) == 0 {
		return es
	}
	es.TitleFilter = filter
	return es
}

func (es *EventStream) FilterContent(filter []string) *EventStream {
	if len(filter) == 0 {
		return es
	}
	es.ContentFilter = filter
	return es
}

func (es *EventStream) Close() *[]Event {
	db.Select(&es.events, es.buildQuery())

	return &es.events
}

func (es EventStream) buildQuery() string {
	var parts []string

	if es.After != 0 {
		parts = append(parts, inizioField+`>`+utils.I64toa(es.After))
	}
	if es.Before != 0 {
		parts = append(parts, fineField+`<`+utils.I64toa(es.Before))
	}
	if len(es.ContentFilter) != 0 {
		var sub string
		for _, f := range es.ContentFilter[:len(es.ContentFilter)-1] {
			sub += contentField + ` like "%` + f + `%" and `
		}
		sub += contentField + ` like "%` + es.ContentFilter[len(es.ContentFilter)-1] + `%"`
		parts = append(parts, sub)
	}
	if len(es.TitleFilter) != 0 {
		var sub string
		for _, f := range es.TitleFilter[:len(es.TitleFilter)-1] {
			sub += titleField + ` like "%` + f + `%" and `
		}
		sub += titleField + ` like "%` + es.TitleFilter[len(es.TitleFilter)-1] + `%"`
		parts = append(parts, sub)
	}

	return baseQuery + strings.Join(parts, " and ")
}
