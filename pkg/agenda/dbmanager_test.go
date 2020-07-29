package agenda

import (
	"fmt"
	"testing"
	"time"

	"github.com/Baldomo/webapi-dav/pkg/utils"
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

	//expected :=
	//	baseQuery +
	//		fmt.Sprintf(`%s>%s and `, inizioField, utils.I64toa(unixPast)) +
	//		fmt.Sprintf(`%s<%s and `, fineField, utils.I64toa(unixNow)) +
	//		fmt.Sprintf(`%s like "%%%s%%" and `, contentField, testStream.ContentFilter[0]) +
	//		fmt.Sprintf(`%s like "%%%s%%" and `, titleField, testStream.TitleFilter[0]) +
	//		fmt.Sprintf(`%s like "%%%s%%"`, titleField, testStream.TitleFilter[1])

	expectedSql := "SELECT summary, description, dtstart, dtend FROM sitoliceo.npjmx_jevents_vevdetail WHERE dtstart > ? AND dtend < ? AND description LIKE ? AND summary LIKE ? AND summary LIKE ?"
	expectedArgs := []interface{}{
		utils.I64toa(unixPast),
		utils.I64toa(unixNow),
		fmt.Sprint("%", testStream.ContentFilter[0], "%"),
		fmt.Sprint("%", testStream.TitleFilter[0], "%"),
		fmt.Sprint("%", testStream.TitleFilter[1], "%"),
	}

	sql, args := testStream.buildQuery()

	if expectedSql != sql {
		t.Errorf("\nExpected SQL: %s\nGot: %s", expectedSql, sql)
	}

	if retCode := checkArgs(expectedArgs, args); retCode > -1 {
		t.Errorf("\nExpected args[%d]: %v\nGot: %v", retCode, expectedArgs[retCode], args[retCode])
	} else if retCode == -2 {
		t.Errorf("\nLength differs: %d and %d\nExpected: %v\nGot: %v", len(expectedArgs), len(args), expectedArgs, args)
	}
}

// Restituisce -1 se nessun errore, -2 se errore lunghezza diversa oppure `i` (indice degli elementi differenti)
func checkArgs(args1, args2 []interface{}) int {
	if len(args1) != len(args2) {
		return -2
	}

	for i, v := range args1 {
		if v != args2[i] {
			return i
		}
	}

	return -1
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
