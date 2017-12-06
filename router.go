package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"crypto/tls"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = EventLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

func NewServer() *http.Server {
	return &http.Server{
		Handler:           NewRouter(),
		Addr:              GetConfig().Conn.Port,
		WriteTimeout:      time.Second * 5,
		ReadTimeout:       time.Second * 5,
		IdleTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
	}
}

func NewServerHTTPS() *http.Server {
	return &http.Server{
		Handler:           NewRouter(),
		Addr:              ":443",
		WriteTimeout:      time.Second * 5,
		ReadTimeout:       time.Second * 5,
		IdleTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		},
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
}
