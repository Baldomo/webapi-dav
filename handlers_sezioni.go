package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	type all struct {
		Genitori Comunicati `json:"genitori" xml:"genitori"`
		Studenti Comunicati `json:"studenti" xml:"studenti"`
		Docenti  Comunicati `json:"docenti" xml:"docenti"`
	}

	switch RequestMime(r.Header) {

	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(all{Genitori, Studenti, Docenti}); err != nil {
			panic(err)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := xml.NewEncoder(w).Encode(all{Genitori, Studenti, Docenti}); err != nil {
			panic(err)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(all{Genitori, Studenti, Docenti}); err != nil {
			panic(err)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
	return
}

func GenitoriComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	var localCount int
	vars := mux.Vars(r)
	count, countPresent := vars["count"]
	if !countPresent {
		localCount = GetLenByName("genitori")
	} else {
		localCount, _ = strconv.Atoi(count)
		if localCount > GetLenByName("genitori") {
			localCount = GetLenByName("genitori")
		}
	}

	switch RequestMime(r.Header) {

	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(GetByName("genitori")[:localCount]); err != nil {
			panic(err)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := xml.NewEncoder(w).Encode(GetByName("genitori")[:localCount]); err != nil {
			panic(err)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(GetByName("genitori")[:localCount]); err != nil {
			panic(err)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func StudentiComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	var localCount int
	vars := mux.Vars(r)
	count, countPresent := vars["count"]
	if !countPresent {
		localCount = GetLenByName("studenti")
	} else {
		localCount, _ = strconv.Atoi(count)
		if localCount > GetLenByName("studenti") {
			localCount = GetLenByName("studenti")
		}
	}

	switch RequestMime(r.Header) {

	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(GetByName("studenti")[:localCount]); err != nil {
			panic(err)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := xml.NewEncoder(w).Encode(GetByName("studenti")[:localCount]); err != nil {
			panic(err)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(GetByName("studenti")[:localCount]); err != nil {
			panic(err)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func DocentiComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	var localCount int
	vars := mux.Vars(r)
	count, countPresent := vars["count"]
	if !countPresent {
		localCount = GetLenByName("docenti")
	} else {
		localCount, _ = strconv.Atoi(count)
		if localCount > GetLenByName("docenti") {
			localCount = GetLenByName("docenti")
		}
	}

	switch RequestMime(r.Header) {

	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(GetByName("docenti")[:localCount]); err != nil {
			panic(err)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := xml.NewEncoder(w).Encode(GetByName("docenti")[:localCount]); err != nil {
			panic(err)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(GetByName("docenti")[:localCount]); err != nil {
			panic(err)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func ProgettiHandler(w http.ResponseWriter, r *http.Request) {
	//TODO prende da file toml/json tutto il necessario, ARRAY di oggetti
	type progetto struct {
		Nome         string `json:"nome" toml:"nome"`
		PathImmagine string `json:"immagine" toml:"immagine"`
		Descr        string `json:"descrizione" toml:"descrizione"`
	}
	w.WriteHeader(http.StatusOK)
}
