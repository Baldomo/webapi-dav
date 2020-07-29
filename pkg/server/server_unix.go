// +build linux darwin

package server

import (
	"net/http"
	"time"

	"github.com/Baldomo/webapi-dav/pkg/config"
	"github.com/Baldomo/webapi-dav/pkg/log"
	"github.com/gorilla/mux"
	"github.com/pseidemann/finish"
)

// Controller interno di un http.Server
type serverHandler struct {
	http *http.Server
}

var (
	timeout = 15 * time.Second

	handler = new(serverHandler)
)

// Inizializza un nuovo router gorilla/mux con sane impostazioni predefinite,
// insieme ai vari middleware
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = log.EventLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}

// Inizializza un http.Server con sani valori di default
func newServer() *http.Server {
	return &http.Server{
		Handler:           NewRouter(),
		Addr:              config.GetConfig().HTTP.Port,
		WriteTimeout:      time.Second * 5,
		ReadTimeout:       time.Second * 5,
		IdleTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
	}
}

//func newServerHTTPS() *http.Server {
//	return &http.Server{
//		Handler:           NewRouter(),
//		Addr:              config.GetConfig().HTTPS.Port,
//		WriteTimeout:      time.Second * 5,
//		ReadTimeout:       time.Second * 5,
//		IdleTimeout:       time.Second * 5,
//		ReadHeaderTimeout: time.Second * 5,
//		TLSConfig: &tls.Config{
//			MinVersion:               tls.VersionTLS12,
//			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
//			PreferServerCipherSuites: true,
//			CipherSuites: []uint16{
//				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
//				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
//				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
//				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
//			},
//		},
//		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
//	}
//}

// Avvia il loop ListenAndServe dell'oggetto server interno (http.Server.ListenAndServe)
func Start() { handler.Start() }
func (sh *serverHandler) Start() {
	sh.http = newServer()

	fin := finish.New()
	fin.Add(sh.http, finish.WithTimeout(timeout))

	go sh.http.ListenAndServe()
	// Blocking: finisher aspetta la conclusione delle richieste
	fin.Wait()
}
