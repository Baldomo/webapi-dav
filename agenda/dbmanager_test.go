package agenda

import (
	"testing"
	"time"
)

var (
	testMalformedEvents = []Event{
		{
			Title:   "Complete event",
			Content: "Content",
			Date: Date{
				Inizio: time.Now().AddDate(0, 0, -1),
				Fine:   time.Now(),
			},
		}, {
			Title: "Missing Content event",
			Date: Date{
				Inizio: time.Now().AddDate(0, 0, -1),
				Fine:   time.Now(),
			},
		}, {
			Title:   "Missing Date event",
			Content: "Content",
		},
	}

	testFullEvents = []Event{
		{
			Title:   "Complete event",
			Content: "Content",
			Date: Date{
				Inizio: time.Now().AddDate(0, 0, -1),
				Fine:   time.Now(),
			},
		}, {
			Title:   "Complete event #2",
			Content: "Content",
			Date: Date{
				Inizio: time.Now().AddDate(0, -1, -1),
				Fine:   time.Now(),
			},
		}, {
			Title:   "Complete event #3",
			Content: "Content",
			Date: Date{
				Inizio: time.Now().AddDate(0, -2, -1),
				Fine:   time.Now(),
			},
		}, {
			Title:   "Complete event #4",
			Content: "Content",
			Date: Date{
				Inizio: time.Now().AddDate(-1, 0, -1),
				Fine:   time.Now(),
			},
		},
	}
)

func TestEvent_FillEmptyFields(t *testing.T) {

	exp := []string{
		"",
		"select description from npjmx_jevents_vevdetail where summary=:summary and dtstart=:dtstart and dtend=:dtend",
		"select dtstart,dtend from npjmx_jevents_vevdetail where summary=:summary and description=:description",
	}

	for i, test := range testMalformedEvents {
		query := buildQuery(&test)
		if exp[i] != query {
			t.Errorf("Expected: %s\nGot: %s", exp[i], query)
		}
	}
}

func TestEventStream_Close(t *testing.T) {
	exp := []string{
		"",
	}
}
