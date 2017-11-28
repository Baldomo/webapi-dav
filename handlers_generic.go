package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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
		absPath, _ := filepath.Abs(GetConfig().General.IndexHTML)
		raw, _ := ioutil.ReadFile(absPath)
		indexHtml = string(raw)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	temp, _ := template.New("index").Parse(indexHtml)
	temp.Execute(w, GetMapOps())
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	var version = struct {
		NomeApi  string `json:"nome" xml:"nome"`
		Versione string `json:"versione" xml:"versione"`
	}{"DaVinci API", Version}
	var versionMessage = APIMessage{http.StatusOK, "DaVinci API v" + Version}

	switch RequestMime(r.Header) {

	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(version); err != nil {
			w.WriteHeader(http.StatusNoContent)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := xml.NewEncoder(w).Encode(version); err != nil {
			w.WriteHeader(http.StatusNoContent)
		}

	case "text/html":
		if err := ShowGenericTemplate(w, versionMessage); err != nil {
			Log.Error(err.Error())
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	var about = struct {
		Autore    string `json:"autore" xml:"autore"`
		Versione  string `json:"versione" xml:"versione"`
		Info      string `json:"info" xml:"info"`
		Copyright string `json:"copyright" xml:"copyright"`
	}{
		"Leonardo Baldin",
		`v` + Version,
		"",
		"(c) 2017",
	}
	var aboutMessage = APIMessage{http.StatusOK, "Leonardo Baldin, v" + Version + ", (c) 2017"}

	switch RequestMime(r.Header) {

	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(about); err != nil {
			w.WriteHeader(http.StatusNoContent)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := xml.NewEncoder(w).Encode(about); err != nil {
			w.WriteHeader(http.StatusNoContent)
		}

	case "text/html":
		if err := ShowGenericTemplate(w, aboutMessage); err != nil {
			Log.Error(err.Error())
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TeapotHandler(w http.ResponseWriter, r *http.Request) {
	var message = APIMessage{http.StatusTeapot, `I'm a teapot`}
	switch RequestMime(r.Header) {

	case "application/json":
		if err := json.NewEncoder(w).Encode(message); err != nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusTeapot)
		}

	case "application/xml":
		if err := xml.NewEncoder(w).Encode(message); err != nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusTeapot)
		}

	case "text/plain":
		fmt.Fprint(w, "I'm a teapot")

	case "text/html":
		if err := ShowGenericTemplate(w, message); err != nil {
			Log.Error(err.Error())
		}

	default:
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
			w.WriteHeader(http.StatusNoContent)
		}

	case "application/xml":
		if err := xml.NewEncoder(w).Encode(message); err != nil {
			w.WriteHeader(http.StatusNoContent)
		}

	}
}
