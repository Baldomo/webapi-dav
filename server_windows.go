package main

import (
	"context"
	"crypto/tls"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type ServerHandler struct {
	Started chan struct{}
	Closing chan struct{}

	http  *http.Server
	https *http.Server
	ipc   *rpc.Server
}

var (
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

func (sh *ServerHandler) Start() {
	sh.Started = make(chan struct{}, 1)
	defer close(sh.Started)
	if GetConfig().HTTPS.Enabled {
		sh.startHTTPS()
	}

	if GetConfig().HTTP.Enabled {
		sh.startHTTP()
	}

	sh.ipc = rpc.NewServer()
	sh.ipc.RegisterName("ServerHandler", sh)
	l, err := net.Listen("tcp", ":2202")
	if err != nil {
		Log.Critical("Impossibile avviare servizio IPC")
	}
	sh.ipc.Accept(l)
	sh.Started <- struct{}{}
}

func (sh *ServerHandler) startHTTP() {
	sh.http = NewServer()
	go func() {
		if err := sh.http.ListenAndServe(); err != nil {
			Log.Fatal(err)
		}
	}()
	return
}

func (sh *ServerHandler) startHTTPS() {
	sh.https = NewServerHTTPS()
	go func() {
		if err := sh.https.ListenAndServeTLS(GetConfig().HTTPS.Cert, GetConfig().HTTPS.Key); err != nil {
			Log.Fatal(err)
		}
	}()
	return
}

func (sh *ServerHandler) restart(_, _ *struct{}) error {
	err := sh.Shutdown(&struct{}{}, &struct{}{})
	if err != nil {
		return err
	}
	sh.Start()
	return nil
}

func (sh *ServerHandler) Shutdown(_, _ *struct{}) error {
	sh.Closing = make(chan struct{}, 1)
	sh.Closing <- struct{}{}
	close(sh.Closing)
	errHttp := shutdown(sh.http)
	if errHttp != nil {
		Log.Error(errHttp.Error())
		return errHttp
	}

	errHttps := shutdown(sh.https)
	if errHttps != nil {
		Log.Error(errHttps.Error())
		return errHttps
	}

	return nil
}

func shutdown(s *http.Server) error {
	if s == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	Log.Warningf("Conclusione richieste con timeout %s", timeout)

	if err := s.Shutdown(ctx); err != nil {
		Log.Error(err.Error())
	} else {
		if s != nil {
			Log.Info("Concluse richieste in arrivo")

			select {
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					return err
				}
			default:
				if deadline, ok := ctx.Deadline(); ok {
					secs := (time.Until(deadline) + time.Second/2) / time.Second
					Log.Warningf("Spegnimento server con timeout %vs", secs)
				}

				done := make(chan error)

				go func() {
					<-ctx.Done()
				}()

				if err := <-done; err != nil {
					return err
				}
			}
		}
	}

	if deadline, ok := ctx.Deadline(); ok {
		secs := (time.Until(deadline) + time.Second/2) / time.Second
		Log.Warningf("Completato spegnimento in %vs", secs)
	}

	return nil
}
