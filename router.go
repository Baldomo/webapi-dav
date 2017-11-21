package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
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
