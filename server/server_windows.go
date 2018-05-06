package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gorilla/mux"
	"leonardobaldin/webapi-dav/config"
	"leonardobaldin/webapi-dav/log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type serverHandler struct {
	Started chan struct{}
	Stopped chan struct{}

	http  *http.Server
	https *http.Server
	ipc   *rpc.Server

	onShutdown []func()
}

var (
	timeout = 15 * time.Second

	handler = new(serverHandler)
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

func newServerHTTPS() *http.Server {
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

func Start() { handler.Start() }
func (sh *serverHandler) Start() {
	sh.Started = make(chan struct{}, 1)
	defer close(sh.Started)
	if config.GetConfig().HTTPS.Enabled {
		sh.startHTTPS()
	}

	if config.GetConfig().HTTP.Enabled {
		sh.startHTTP()
	}

	sh.ipc = rpc.NewServer()
	sh.ipc.RegisterName("serverHandler", sh)
	l, err := net.Listen("tcp", ":2202")
	if err != nil {
		log.Log.Critical("Impossibile avviare servizio IPC")
	}
	sh.ipc.Accept(l)
	sh.Started <- struct{}{}
}

func (sh *serverHandler) startHTTP() {
	sh.http = newServer()
	go func() {
		if err := sh.http.ListenAndServe(); err != nil {
			log.Log.Fatal(err)
		}
	}()
	return
}

func (sh *serverHandler) startHTTPS() {
	sh.https = newServerHTTPS()
	go func() {
		if err := sh.https.ListenAndServeTLS(config.GetConfig().HTTPS.Cert, config.GetConfig().HTTPS.Key); err != nil {
			log.Log.Fatal(err)
		}
	}()
	return
}

func (sh *serverHandler) restart(_, _ *struct{}) error {
	err := sh.Shutdown(&struct{}{}, &struct{}{})
	if err != nil {
		return err
	}
	sh.Start()
	return nil
}

func OnShutdown(f func()) {
	handler.onShutdown = append(handler.onShutdown, f)
}

func Shutdown() { handler.Shutdown(&struct{}{}, &struct{}{}) }
func (sh *serverHandler) Shutdown(_, _ *struct{}) error {
	sh.Stopped = make(chan struct{}, 2)

	for _, f := range sh.onShutdown {
		f()
	}

	errHttp := shutdown(sh.http, sh.Stopped)
	if errHttp != nil {
		log.Log.Error(errHttp.Error())
		return errHttp
	}

	errHttps := shutdown(sh.https, sh.Stopped)
	if errHttps != nil {
		log.Log.Error(errHttps.Error())
		return errHttps
	}

	select {
	case sh.Stopped <- struct{}{}:
		close(sh.Stopped)
		return fmt.Errorf("impossibile terminare server automaticamente")
	default:
		break
	}

	close(sh.Stopped)
	return nil
}

func shutdown(s *http.Server, cchan chan struct{}) error {
	if s == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Log.Warningf("Conclusione richieste con timeout %s", timeout)

	if err := s.Shutdown(ctx); err != nil {
		log.Log.Error(err.Error())
	} else {
		if s != nil {
			log.Log.Info("Concluse richieste in arrivo")

			select {
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					return err
				}
			default:
				if deadline, ok := ctx.Deadline(); ok {
					secs := (time.Until(deadline) + time.Second/2) / time.Second
					log.Log.Warningf("Spegnimento server con timeout %vs", secs)
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
		cchan <- struct{}{}
		secs := (time.Until(deadline) + time.Second/2) / time.Second
		log.Log.Warningf("Completato spegnimento in %vs", secs)
	}

	log.CloseLogger()

	return nil
}
