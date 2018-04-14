package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServers(t *testing.T) {
	/*tests := []struct {
		desc         string
		endpoint     string
		expectedCode int
		method       string
	}{
		{
			"Endpoint classi",
			"/classi",
			http.StatusOK,
			"GET",
		}, {
			"Endpoint comunicati",
			"/comunicati",
			http.StatusOK,
			"GET",
		}, {
			"Endpoints comuncati docenti",
			"/comunicati/docenti",
			http.StatusOK,
			"GET",
		}, {
			"Endpoint comunicati genitori",
			"/comunicati/genitori",
			http.StatusOK,
			"GET",
		}, {
			"Endpoint comunicati studenti",
			"/comunicati/studenti",
			http.StatusOK,
			"GET",
		}, {
			"Endpoint docenti",
			"/docenti",
			http.StatusOK,
			"GET",
		},
	}*/

	/*go Handler.Start()
	<-Handler.Started

	for _, test := range tests {
		req, _ := http.NewRequest(test.method, test.endpoint, nil)
		resp := executeRequest(req)
		checkResponseCode(t, test.expectedCode, resp.Code)
	}*/
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	Handler.https.Handler.ServeHTTP(rr, req)
	Handler.http.Handler.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
