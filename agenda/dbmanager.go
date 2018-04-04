package agenda

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"time"
	"leonardobaldin/webapi-dav/utils"
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

	baseQuery = "select " + titleField + "," + contentField + "," + inizioField + "," + fineField +
		" from " + agendaTable + " where "
)

var (
	db *sqlx.DB
)

type EventStream struct {
	after         int64
	before        int64
	titleFilter   []string
	contentFilter []string
	events        []Event
}

type Date struct {
	Inizio time.Time `json:"inizio" xml:"Inizio"`
	Fine   time.Time `json:"fine" xml:"Fine"`
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

func GetInInterval(start time.Time, end time.Time) (*[]Event, error) {
	var e []Event
	startUnix := strconv.Itoa(int(start.Unix()))
	endUnix := strconv.Itoa(int(end.Unix()))
	err := db.Select(&e,
		"select "+titleField+","+contentField+","+inizioField+","+fineField+
			" from "+agendaTable+
			" where "+inizioField+">="+startUnix+" and "+fineField+"<="+endUnix)
	return &e, err
}

func NewEventStream() *EventStream {
	return &EventStream{
		after:  0,
		before: 0,
		titleFilter: []string{},
		contentFilter: []string{},
		events: []Event{},
	}
}

func (es *EventStream) GetAfter(epoch int64) *EventStream {
	if epoch > es.before || epoch == 0 {
		return es
	}
	es.after = epoch
	return es
}

func (es *EventStream) GetBefore(epoch int64) *EventStream {
	if epoch < es.after {
		return es
	}
	es.before = epoch
	return es
}

func (es *EventStream) FilterTitle(filter []string) *EventStream {
	if len(filter) == 0 {
		return es
	}
	es.titleFilter = filter
	return es
}

func (es *EventStream) FilterContent(filter []string) *EventStream {
	if len(filter) == 0 {
		return es
	}
	es.contentFilter = filter
	return es
}

func (es *EventStream) Close() *[]Event {
	var query = baseQuery
	if es.after != 0 {
		query += inizioField + `>` + utils.I64toa(es.after)
	}
	if es.before != 0 {
		if es.after != 0 {
			query += " and "
		}
		query += fineField + `<` + utils.I64toa(es.before)
	}
	if len(es.contentFilter) != 0 {
		for _, f := range es.contentFilter[:len(es.contentFilter)-1] {
			query += contentField + ` like "%` + f + `%" and `
		}
		query += contentField + ` like "%` + es.contentFilter[len(es.contentFilter)-1] + `%"`
	}
	if len(es.titleFilter) != 0 {
		if len(es.contentFilter) != 0 {
			query += " and "
		}
		for _, f := range es.titleFilter[:len(es.titleFilter)-1] {
			query += titleField + ` like "%` + f + `%" and `
		}
		query += titleField + ` like "%` + es.titleFilter[len(es.titleFilter)-1] + `%"`
	}

	db.Select(&es.events, query)

	return &es.events
}
