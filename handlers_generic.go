package main

import (
	"encoding/json"
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

var (
	indexHtml = ""
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if indexHtml == "" {
		absPath, _ := filepath.Abs(filepath.Join(GetConfig().Dirs.HTML, "index.html"))
		raw, _ := ioutil.ReadFile(absPath)
		indexHtml = string(raw)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	temp, _ := template.New("index").Parse(indexHtml)
	temp.Execute(w, GetMapOps())
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	var versionMessage = APIMessage{http.StatusOK, "webapi-dav v" + VersionNumber}

	w.WriteHeader(http.StatusOK)
	switch RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(versionMessage); err != nil {
			Log.Error("VersionHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(versionMessage); err != nil {
			Log.Error("VersionHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "text/html":
		if err := ShowGenericTemplate(w, versionMessage); err != nil {
			Log.Error("VersionHandler: errore template html")
		}

	default:
		Log.Info("VersionHandler: Accept errato")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	var aboutMessage = APIMessage{http.StatusOK, "Leonardo Baldin, v" + VersionNumber + ", (c) 2017"}

	w.WriteHeader(http.StatusOK)
	switch RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(aboutMessage); err != nil {
			Log.Error("AboutHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(aboutMessage); err != nil {
			Log.Error("AboutHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "text/html":
		if err := ShowGenericTemplate(w, aboutMessage); err != nil {
			Log.Error("AboutHandler: errore template html")
		}

	default:
		Log.Warning("AboutHandler: Accept errato")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TeapotHandler(w http.ResponseWriter, r *http.Request) {
	var message = APIMessage{http.StatusTeapot, `I'm a teapot`}

	w.WriteHeader(http.StatusTeapot)
	switch RequestMime(r.Header) {
	case "application/json":
		if err := json.NewEncoder(w).Encode(message); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "application/xml":
		if err := xml.NewEncoder(w).Encode(message); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "text/html":
		if err := ShowGenericTemplate(w, message); err != nil {
			Log.Error("TeapotHandler: errore template html")
		}

	default:
		Log.Warning("TeapotHandler: Accept errato")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	var message = APIMessage{http.StatusNotFound, "Non trovato"}

	w.WriteHeader(http.StatusNotFound)
	switch RequestMime(r.Header) {
	case "text/html":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := ShowGenericTemplate(w, message); err != nil {
			Log.Error(err.Error())
		}
	case "application/json":
		if err := json.NewEncoder(w).Encode(message); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "application/xml":
		if err := xml.NewEncoder(w).Encode(message); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		Log.Warning("NotFoundHandler: Accept errato")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func RefreshHTML() {
	Log.Info("Ricaricamento pagine web...")
	indexHtml = ""
	absPath, _ := filepath.Abs(filepath.Join(GetConfig().Dirs.HTML, "index.html"))
	raw, _ := ioutil.ReadFile(absPath)
	indexHtml = string(raw)
}
