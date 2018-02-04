// +build linux darwin

package server

import (
	"context"
	"crypto/tls"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"leonardobaldin/webapi-dav/config"
	"leonardobaldin/webapi-dav/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type serverHandler struct {
	Started chan struct{}
	Closing chan struct{}

	http  *http.Server
	https *http.Server
}

var (
	signals chan os.Signal
	timeout = 15 * time.Second

	Handler = new(serverHandler)
)

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

func NewServer() *http.Server {
	return &http.Server{
		Handler:           NewRouter(),
		Addr:              config.GetConfig().HTTP.Port,
		WriteTimeout:      time.Second * 5,
		ReadTimeout:       time.Second * 5,
		IdleTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
	}
}

func NewServerHTTPS() *http.Server {
	return &http.Server{
		Handler:           NewRouter(),
		Addr:              config.GetConfig().HTTPS.Port,
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

func (sh *serverHandler) Start() {
	sh.Started = make(chan struct{}, 1)
	defer close(sh.Started)
	if config.GetConfig().HTTPS.Enabled {
		sh.startHTTPS()
	}

	if config.GetConfig().HTTP.Enabled {
		sh.startHTTP()
	}

	sh.Started <- struct{}{}
}

func (sh *serverHandler) startHTTP() {
	sh.http = NewServer()
	go func() {
		if err := sh.http.ListenAndServe(); err != nil {
			//Log.Fatal(err)
		}
	}()
	Shutdown(sh.http)
	return
}

func (sh *serverHandler) startHTTPS() {
	sh.https = NewServerHTTPS()
	go func() {
		if err := sh.https.ListenAndServeTLS(config.GetConfig().HTTPS.Cert, config.GetConfig().HTTPS.Key); err != nil {
			//Log.Fatal(err)
		}
	}()
	Shutdown(sh.https)
	return
}

func Shutdown(s *http.Server) {
	signals = make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals

	shutdown(s, log.Log)
}

func shutdown(s *http.Server, logger *logging.Logger) {
	if s == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Warningf("Conclusione richieste con timeout %s", timeout)

	if err := s.Shutdown(ctx); err != nil {
		logger.Error(err.Error())
	} else {
		if s != nil {
			logger.Info("Concluse richieste in arrivo")

			select {
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					logger.Error(err.Error())
					return
				}
			default:
				if deadline, ok := ctx.Deadline(); ok {
					secs := (time.Until(deadline) + time.Second/2) / time.Second
					logger.Warningf("Spegnimento server con timeout %vs", secs)
				}

				done := make(chan error)

				go func() {
					<-ctx.Done()
				}()

				if err := <-done; err != nil {
					logger.Error(err.Error())
					return
				}
			}
		}
	}

	if deadline, ok := ctx.Deadline(); ok {
		secs := (time.Until(deadline) + time.Second/2) / time.Second
		logger.Warningf("Completato spegnimento in %vs", secs)
	}
}
