package server

import (
	"encoding/json"
	"strings"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/Baldomo/webapi-dav/pkg/config"
	"github.com/Baldomo/webapi-dav/pkg/log"
	"github.com/Baldomo/webapi-dav/pkg/utils"
	"github.com/Baldomo/webapi-dav/pkg/auth"
)

var (
	indexHtml = ""
)

// Middleware per la verifica del token JWT di autorizzazione
func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Divide l'header
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		// Verifica l'header
		if len(authHeader) != 2 {
			log.Error("AuthorizationMiddleware: header is not valid")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed or missing bearer token"))

		} else {
			// Ottiene il token
			log.Error(authHeader[1])
			jwtToken := []byte(authHeader[1])

			// Verifica il token e passa la richiesta all'handler se valido
			if _, err := auth.ParseToken(jwtToken); err != nil {
				log.Error(err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if indexHtml == "" {
		absPath, _ := filepath.Abs(filepath.Join(config.GetConfig().Dirs.HTML, "index.html"))
		raw, _ := ioutil.ReadFile(absPath)
		indexHtml = string(raw)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp, _ := template.New("index").Parse(indexHtml)
	err := temp.Execute(w, utils.TemplateData())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func OpenapiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/yaml")
	http.ServeFile(w, r, filepath.Join(config.GetConfig().Dirs.HTML, "openapi.yaml"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	var aboutMessage = utils.APIMessage{
		Code: http.StatusOK,
		Info: "Leonardo Baldin, v" + utils.VersionNumber + ", (c) 2017",
	}

	defer r.Body.Close()
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(aboutMessage); err != nil {
			log.Error("AboutHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		if err := utils.ShowGenericTemplate(w, aboutMessage); err != nil {
			log.Error("AboutHandler: errore template html")
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	var message = utils.APIMessage{Code: http.StatusNotFound, Info: "Non trovato"}

	defer r.Body.Close()
	w.WriteHeader(http.StatusNotFound)
	switch utils.RequestMime(r.Header) {
	case "text/html":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := utils.ShowGenericTemplate(w, message); err != nil {
			log.Error(err.Error())
		}

	case "application/json":
		if err := json.NewEncoder(w).Encode(message); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func TeapotHandler(w http.ResponseWriter, r *http.Request) {
	var message = utils.APIMessage{Code: http.StatusTeapot, Info: `I'm a teapot`}

	defer r.Body.Close()
	w.WriteHeader(http.StatusTeapot)
	switch utils.RequestMime(r.Header) {
	case "application/json":
		if err := json.NewEncoder(w).Encode(message); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusTeapot)
		}

	case "text/html":
		if err := utils.ShowGenericTemplate(w, message); err != nil {
			log.Error("TeapotHandler: errore template html")
		} else {
			w.WriteHeader(http.StatusTeapot)
		}
	}
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	var versionMessage = utils.APIMessage{Code: http.StatusOK, Info: "webapi-dav v" + utils.VersionNumber}

	defer r.Body.Close()
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(versionMessage); err != nil {
			log.Error("VersionHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		if err := utils.ShowGenericTemplate(w, versionMessage); err != nil {
			log.Error("VersionHandler: errore template html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

// Ricarica in memoria le pagine HTML statiche (che verranno servite aggiornate
// alle richieste successive)
func RefreshHTML() {
	log.Info("Ricaricamento pagine web...")
	indexHtml = ""
	absPath, _ := filepath.Abs(filepath.Join(config.GetConfig().Dirs.HTML, "index.html"))
	raw, _ := ioutil.ReadFile(absPath)
	indexHtml = string(raw)
}
