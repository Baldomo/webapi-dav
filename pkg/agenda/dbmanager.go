package agenda

import (
	"fmt"
	"os"

	"github.com/Baldomo/webapi-dav/pkg/config"
	"github.com/Baldomo/webapi-dav/pkg/log"
	"github.com/Baldomo/webapi-dav/pkg/utils"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	inizioField  = "dtstart"
	fineField    = "dtend"
	contentField = "description"
	titleField   = "summary"
)

var (
	db *sqlx.DB

	agendaTable = config.GetConfig().DB.Schema + ".npjmx_jevents_vevdetail"
)

// Definisce un modello di range di eventi direttamente decodificato
// da una richiesta al relativo endpoint
type EventStream struct {
	// Data di inizio del range
	After int64 `json:"dopo,omitempty"`

	// Data di fine del range
	Before int64 `json:"prima,omitempty"`

	// Filtro stringa per il titolo (case insensitive), può essere varie stringhe
	// che vengono controllate separatamente
	TitleFilter []string `json:"filtri_titolo,omitempty"`

	// Filtro stringa per il titolo (case insensitive), può essere varie stringhe
	// che vengono controllate separatamente
	ContentFilter []string `json:"filtri_contenuto,omitempty"`

	events []Event
}

// Rappresenta un singolo evento come contenuto nel database Joomla
type Event struct {
	// Inizio dell'evento (in Unix epoch)
	Inizio int64 `json:"inizio" db:"dtstart"`

	// Fine dell'evento (in Unix epoch)
	Fine int64 `json:"fine" db:"dtend"`

	// Testo allegato all'evento
	Content string `json:"contenuto" db:"description"`

	// Titolo dell'evento
	Title string `json:"titolo" db:"summary"`
}

// Esegue connessione al database con log di errori eventuali
func Fetch() {
	user := os.Getenv("WEBAPI_DB_USER")
	pwd := os.Getenv("WEBAPI_DB_PWD")
	if user == "" || pwd == "" {
		// Impossibile determinare credenziali per il DB
		return
	}
	dataSource := fmt.Sprint(user, ":", pwd, "@/")

	var err error
	db, err = sqlx.Connect("mysql", dataSource)
	if err != nil {
		log.Critical("Errore collegamento a database")
		log.Critical(err.Error())
	}
}

// Costruttore per EventStream
func NewEventStream() *EventStream {
	return &EventStream{
		After:         0,
		Before:        0,
		TitleFilter:   []string{},
		ContentFilter: []string{},
		events:        []Event{},
	}
}

// Modifica un EventStream cambiando il campo EvenStream.After e ritorna
// un puntatore allo stesso EventStream (builder pattern)
func (es *EventStream) GetAfter(epoch int64) *EventStream {
	if es.Before != 0 && (epoch > es.Before || epoch == 0) {
		return es
	}
	es.After = epoch
	return es
}

// Modifica un EventStream cambiando il campo EvenStream.Before e ritorna
// un puntatore allo stesso EventStream (builder pattern)
func (es *EventStream) GetBefore(epoch int64) *EventStream {
	if epoch < es.After {
		return es
	}
	es.Before = epoch
	return es
}

// Cambia EventStream.TitleFilter
func (es *EventStream) FilterTitle(filter []string) *EventStream {
	if len(filter) == 0 {
		return es
	}
	es.TitleFilter = filter
	return es
}

// Cambia EventStream.ContentFilter
func (es *EventStream) FilterContent(filter []string) *EventStream {
	if len(filter) == 0 {
		return es
	}
	es.ContentFilter = filter
	return es
}

// Chiude un EventStream: compila una query sicura in base ai campi di EventStream
// e restituisce gli eventi ottenuti dal database
func (es *EventStream) Close() (*[]Event, error) {
	if db == nil {
		msg := "Database non inizializzato"
		log.Error(msg)
		return &[]Event{}, fmt.Errorf(msg)
	}

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

// Funzione interna di utility per compilare una query compatibile con MySQL
// in base ai campi di EventStream
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
