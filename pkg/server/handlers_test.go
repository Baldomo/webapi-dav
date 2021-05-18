package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Baldomo/webapi-dav/pkg/auth"
	"github.com/gbrlsnchs/jwt/v3"
)

type customPayload struct {
	Payload jwt.Payload
}

// var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7ImlzcyI6ImRhdiIsInN1YiI6Im5vbWUuY29nbm9tZSIsImF1ZCI6WyJodHRwOi8vZXhhbXBsZS5vcmciLCJodHRwczovL2V4YW1wbGUub3JnIl0sImV4cCI6MjU5MzU0MDQwOSwiaWF0IjoxNTkwOTQ4NDA5fX0.0NARK5Q7Sm_KH_qYPXbu26ihGzYheJHp7-cokhusCEM"

func TestAuthnMiddleware(t *testing.T) {
	handler := http.HandlerFunc(PdfHandler)

	t.Run("richiesta non autorizzata", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/sitoLiceo/images/comunicati/comunicati-docenti/0.pdf", nil)
		if err != nil {
			t.Fatal(err)
		}

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusUnauthorized {
			t.Errorf("handler ha restituito codice errato: avuto %v atteso %v",
				status, http.StatusUnauthorized)
		}
	})

	// secret := config.GetConfig().Auth.JWTSecret
	// fqdn := config.GetConfig().General.FQDN

	jwtSigner := jwt.NewHS256([]byte("secret"))
	aud := jwt.Audience{"http://liceodavinci.tv", "https://liceodavinci.tv"}
	exp := jwt.NumericDate(time.Now().Add(time.Minute))
	iss := jwt.NumericDate(time.Now())
	pl := customPayload{jwt.Payload{
		Audience:       aud,
		ExpirationTime: exp,
		IssuedAt:       iss,
		Subject:        "test",
	}}
	token, _ := jwt.Sign(&pl, jwtSigner)

	t.Run("richiesta autorizzata", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/sitoLiceo/images/comunicati/comunicati-docenti/0.pdf", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+string(token))

		auth.InitializeSigning()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Log(req)
			t.Log(rr.Body.String())
			t.Errorf("handler ha restituito codice errato: avuto %v atteso %v",
				status, http.StatusOK)
		}
	})
}
