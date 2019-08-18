// +build linux darwin

package server

import (
	"context"
	"github.com/Baldomo/webapi-dav/internal/config"
	"github.com/Baldomo/webapi-dav/internal/log"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type serverHandler struct {
	Stopped chan struct{}

	http *http.Server
}

var (
	signals chan os.Signal
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

func Start() { handler.Start() }
func (sh *serverHandler) Start() {
	sh.http = newServer()
	go sh.http.ListenAndServe()

	sh.Stopped = make(chan struct{}, 2)
	signals = make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals

	err := shutdown(sh.http, sh.Stopped)
	if err != nil {
		log.Log.Error(err.Error())
	}

	select {
	case sh.Stopped <- struct{}{}:
		close(sh.Stopped)
	default:
		break
	}

	close(sh.Stopped)
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

	if deadline, ok := ctx.Deadline(); ok {
		cchan <- struct{}{}
		secs := (time.Until(deadline) + time.Second/2) / time.Second
		log.Log.Warningf("Completato spegnimento in %vs", secs)
	}

	log.CloseLogger()

	return nil
}
