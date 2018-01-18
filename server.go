package main

import (
	"context"
	"crypto/tls"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type serverHandler struct {
	servers []*http.Server
	done    chan struct{}
}

var (
	signals chan os.Signal
	timeout = 15 * time.Second
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
		Addr:              GetConfig().HTTP.Port,
		WriteTimeout:      time.Second * 5,
		ReadTimeout:       time.Second * 5,
		IdleTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
	}
}

func NewServerHTTPS() *http.Server {
	return &http.Server{
		Handler:           NewRouter(),
		Addr:              GetConfig().HTTPS.Port,
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

func Shutdown(s *http.Server) {
	signals = make(chan os.Signal, 1)

	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	shutdown(s, Log)
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

func StartServers() {
	/*if GetConfig().Conn.FastCGI {
		router := NewRouter()
		Log.Fatal(fcgi.Serve(nil, router))
	}*/
	if GetConfig().HTTPS.Enabled {
		startHTTPS()
	}
	if GetConfig().HTTP.Enabled {
		startHTTP()
	}
}

func startHTTP() {
	httpServer := NewServer()
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			Log.Fatal(err)
		}
	}()
	Shutdown(httpServer)
	return
}

func startHTTPS() {
	httpsServer := NewServerHTTPS()
	go func() {
		if err := httpsServer.ListenAndServeTLS(GetConfig().HTTPS.Cert, GetConfig().HTTPS.Key); err != nil {
			Log.Fatal(err)
		}
	}()
	Shutdown(httpsServer)
	return
}

func Restart(s *http.Server, name string) {
	Log.Info("Riavvio server %s in corso...", name)
}
