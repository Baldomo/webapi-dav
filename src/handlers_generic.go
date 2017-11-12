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
	// Orrendo ma funzionale
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
	type version struct {
		NomeApi  string `json:"nome" xml:"nome"`
		Versione string `json:"versione" xml:"versione"`
	}

	switch RequestMime(r.Header) {

	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(version{"API Dav", "1.0"}); err != nil {
			w.WriteHeader(http.StatusNoContent)
		}

	case "application/xml":

		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := xml.NewEncoder(w).Encode(version{"API Dav", "1.0"}); err != nil {
			w.WriteHeader(http.StatusNoContent)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	type about struct {
		Autore    string `json:"autore" xml:"autore"`
		Versione  string `json:"versione" xml:"versione"`
		Info      string `json:"info" xml:"info"`
		Copyright string `json:"copyright" xml:"copyright"`
	}
	myAbout := about{
		"Leonardo Baldin",
		"v1.0",
		"",
		"2017",
	}
	switch RequestMime(r.Header) {

	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(myAbout); err != nil {
			w.WriteHeader(http.StatusNoContent)
		}
	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := xml.NewEncoder(w).Encode(myAbout); err != nil {
			w.WriteHeader(http.StatusNoContent)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func TeapotHandler(w http.ResponseWriter, r *http.Request) {
	switch RequestMime(r.Header) {

	case "application/json":
		if err := json.NewEncoder(w).Encode("I'm a teapot"); err != nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusTeapot)
		}

	case "application/xml":
		if err := xml.NewEncoder(w).Encode("I'm a teapot"); err != nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusTeapot)
		}

	case "text/plain":
		fmt.Fprint(w, "I'm a teapot")

	default:
		w.WriteHeader(http.StatusBadRequest)

	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {

}
