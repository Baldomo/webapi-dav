package agenda

import (
	"fmt"

	"github.com/Baldomo/webapi-dav/internal/config"
	"github.com/Baldomo/webapi-dav/internal/log"
	"github.com/Baldomo/webapi-dav/internal/utils"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
		log.Critical("Errore collegamento a database")
		log.Critical(err.Error())
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

func (es *EventStream) Close() (*[]Event, error) {
	rows, err := db.Query(es.buildQuery())
	if err != nil {
		log.Error(err.Error())
		return &es.events, err
	}
	defer rows.Close()

	for rows.Next() {
		e := Event{}
		err = rows.Scan(&e.Title, &e.Content, &e.Inizio, &e.Fine)
		if err != nil {
			log.Error(err.Error())
			return &es.events, err
		}
		es.events = append(es.events, e)
	}

	return &es.events, nil
}

func (es EventStream) buildQuery() (string, []interface{}) {
	query := sq.Select(titleField, contentField, inizioField, fineField).
		From(agendaTable)

	if es.After != 0 {
		query = query.Where(sq.Gt{
			inizioField: utils.I64toa(es.After),
		})
	}

	if es.Before != 0 {
		query = query.Where(sq.Lt{
			fineField: utils.I64toa(es.Before),
		})
	}

	if len(es.ContentFilter) > 0 {
		for _, filter := range es.ContentFilter {
			query = query.Where(contentField+" LIKE ?", fmt.Sprint("%", filter, "%"))
		}
	}

	if len(es.TitleFilter) > 0 {
		for _, filter := range es.TitleFilter {
			query = query.Where(titleField+" LIKE ?", fmt.Sprint("%", filter, "%"))
		}
	}

	sql, args, _ := query.ToSql()
	return sql, args
}
