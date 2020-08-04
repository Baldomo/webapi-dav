package server

import (
	"context"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/Baldomo/webapi-dav/pkg/config"
	"github.com/Baldomo/webapi-dav/pkg/log"
	"github.com/gorilla/mux"
)

// Controller virtuale per un oggetto server con IPC
// e controllo di spegnimento tramite un canale
type serverHandler struct {
	Stopped chan struct{}

	http *http.Server
	ipc  *rpc.Server
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

		// Aggiunge il middleware di autorizzazione se richiesto
		if route.RequiresAuthorization {
			handler = AuthorizationMiddleware(handler).(http.Handler)
		}

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
	go sh.http.ListenAndServe()

	sh.Stopped = make(chan struct{}, 2)

	sh.ipc = rpc.NewServer()
	sh.ipc.RegisterName("serverHandler", sh)
	l, err := net.Listen("tcp", ":2202")
	if err != nil {
		log.Critical("Impossibile avviare servizio IPC")
	}

	// Blocking: aspetta segnali da attraverso IPC
	sh.ipc.Accept(l)
}

// Riavvia il server interno
func (sh *serverHandler) restart(_, _ *struct{}) error {
	err := sh.Shutdown(&struct{}{}, &struct{}{})
	if err != nil {
		return err
	}
	sh.Start()
	return nil
}

// Esegue graceful shutdown del server interno (termina tutte le richieste in
// corso, rifiuta quelle in arrivo)
func Shutdown() { handler.Shutdown(&struct{}{}, &struct{}{}) }
func (sh *serverHandler) Shutdown(_, _ *struct{}) error {
	err := shutdown(sh.http, sh.Stopped)
	if err != nil {
		log.Error(err.Error())
	}

	select {
	case sh.Stopped <- struct{}{}:
		close(sh.Stopped)
	default:
		break
	}

	close(sh.Stopped)
	return nil
}

// Metodo interno di controllo dello spegnimento
func shutdown(s *http.Server, cchan chan struct{}) error {
	if s == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Warningf("Conclusione richieste con timeout %s", timeout)

	if err := s.Shutdown(ctx); err != nil {
		log.Error(err.Error())
	} else {
		if s != nil {
			log.Info("Concluse richieste in arrivo")

			select {
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					return err
				}
			default:
				if deadline, ok := ctx.Deadline(); ok {
					secs := (time.Until(deadline) + time.Second/2) / time.Second
					log.Warningf("Spegnimento server con timeout %vs", secs)
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
		log.Warningf("Completato spegnimento in %vs", secs)
	}

	log.CloseLogger()

	return nil
}
