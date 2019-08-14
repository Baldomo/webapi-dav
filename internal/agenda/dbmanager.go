package agenda

import (
	"github.com/Baldomo/webapi-dav/internal/config"
	"github.com/Baldomo/webapi-dav/internal/log"
	"github.com/Baldomo/webapi-dav/internal/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

const (
	dataSource = "apiliceo:apiliceo2018-txc122tr887@/"

	inizioField  = "dtstart"
	fineField    = "dtend"
	contentField = "description"
	titleField   = "summary"
)

var (
	db *sqlx.DB

	agendaTable = config.GetConfig().DB.Schema + ".npjmx_jevents_vevdetail"
	baseQuery = "select " + titleField + "," + contentField + "," + inizioField + "," + fineField +
		" from " + agendaTable + " where "
)

type EventStream struct {
	After         int64    `json:"dopo,omitempty"`
	Before        int64    `json:"prima,omitempty"`
	TitleFilter   []string `json:"filtri_titolo,omitempty"`
	ContentFilter []string `json:"filtri_contenuto,omitempty"`
	events        []Event
}

type Event struct {
	Inizio  int64  `json:"inizio" db:"dtstart"`
	Fine    int64  `json:"fine" db:"dtend"`
	Content string `json:"contenuto" db:"description"`
	Title   string `json:"titolo" db:"summary"`
}

func Fetch() {
	var err error
	db, err = sqlx.Connect("mysql", dataSource)
	if err != nil {
		log.Log.Critical("Errore collegamento a database")
		log.Log.Critical(err.Error())
	}
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
	rows, err := db.Query(es.buildQuery())
	defer rows.Close()
	if err != nil {
		log.Log.Error(err.Error())
	}

	for rows.Next() {
		e := Event{}
		err = rows.Scan(&e.Title, &e.Content, &e.Inizio, &e.Fine)
		if err != nil {
			log.Log.Error(err.Error())
		}
		es.events = append(es.events, e)
	}

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