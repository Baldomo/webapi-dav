package main

import (
	"testing"
	"time"
)

func TestEvent_FillEmptyFields(t *testing.T) {
	tests := []struct {
		expected string
		e        Event
	}{
		{
			"",
			Event{
				Title:   "Complete event",
				Content: "Content",
				Date: Date{
					Inizio: time.Now().AddDate(0, 0, -1),
					Fine:   time.Now(),
				},
			},
		}, {
			"select description from npjmx_jevents_vevdetail where summary=:summary and dtstart=:dtstart and dtend=:dtend",
			Event{
				Title: "Missing Content event",
				Date: Date{
					Inizio: time.Now().AddDate(0, 0, -1),
					Fine:   time.Now(),
				},
			},
		}, {
			"select dtstart,dtend from npjmx_jevents_vevdetail where summary=:summary and description=:description",
			Event{
				Title:   "Missing Date event",
				Content: "Content",
			},
		},
	}

	for _, test := range tests {
		query := buildQuery(&test.e)
		if test.expected != query {
			t.Errorf("Expected: %s\nGot: %s", test.expected, query)
		}
	}
}
