package agenda

import (
	"fmt"
	"github.com/Baldomo/webapi-dav/utils"
	"testing"
	"time"
)

var (
	testMalformedEvents = []Event{
		{
			Title:   "Complete event",
			Content: "Content",
			Inizio:  time.Now().AddDate(0, 0, -1).Unix(),
			Fine:    time.Now().Unix(),
		}, {
			Title:  "Missing Content event",
			Inizio: time.Now().AddDate(0, 0, -1).Unix(),
			Fine:   time.Now().Unix(),
		}, {
			Title:   "Missing Date event",
			Content: "Content",
		}, {
			Title: "Missing Date and Content event",
		},
	}

	testFullEvents = []Event{
		{
			Title:   "Complete event",
			Content: "Content",
			Inizio:  time.Now().AddDate(0, 0, -1).Unix(),
			Fine:    time.Now().Unix(),
		}, {
			Title:   "Complete event #2",
			Content: "Content",
			Inizio:  time.Now().AddDate(0, -1, -1).Unix(),
			Fine:    time.Now().Unix(),
		}, {
			Title:   "Complete event #3",
			Content: "Content",
			Inizio:  time.Now().AddDate(0, -2, -1).Unix(),
			Fine:    time.Now().Unix(),
		}, {
			Title:   "Complete event #4",
			Content: "Content",
			Inizio:  time.Now().AddDate(-1, 0, -1).Unix(),
			Fine:    time.Now().Unix(),
		},
	}
)

func TestEventStream_Close(t *testing.T) {
	unixPast := time.Now().AddDate(0, -6, 0).Unix()
	unixNow := time.Now().Unix()

	var testStream = EventStream{
		After:  unixPast,
		Before: unixNow,
		TitleFilter: []string{
			"aaaaa",
			"bbbbb",
		},
		ContentFilter: []string{
			"ccccc",
		},
	}

	expected :=
		baseQuery +
			fmt.Sprintf(`%s>%s and `, inizioField, utils.I64toa(unixPast)) +
			fmt.Sprintf(`%s<%s and `, fineField, utils.I64toa(unixNow)) +
			fmt.Sprintf(`%s like "%%%s%%" and `, contentField, testStream.ContentFilter[0]) +
			fmt.Sprintf(`%s like "%%%s%%" and `, titleField, testStream.TitleFilter[0]) +
			fmt.Sprintf(`%s like "%%%s%%"`, titleField, testStream.TitleFilter[1])

	output := testStream.buildQuery()

	if expected != output {
		t.Errorf("Expected: %s\nGot: %s", expected, output)
	}
}

func TestEventStream_GetBefore(t *testing.T) {
	expected := &EventStream{
		Before: int64(951195600),
	}
	testTime, _ := time.Parse(time.RFC822Z, "22 Feb 00 05:00 +0000")
	es := NewEventStream().GetBefore(testTime.Unix())

	if expected.Before != es.Before {
		t.Errorf("Expected before %d, got %d", expected.Before, es.Before)
	}
}

func TestEventStream_GetAfter(t *testing.T) {
	expected := &EventStream{
		After: int64(951195600),
	}
	testTime, _ := time.Parse(time.RFC822Z, "22 Feb 00 05:00 +0000")
	es := NewEventStream().GetAfter(testTime.Unix())

	if expected.After != es.After {
		t.Errorf("Expected after %d, got %d", expected.After, es.After)
	}
}