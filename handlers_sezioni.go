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

	w.WriteHeader(http.StatusOK)
	switch RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(all{Genitori, Studenti, Docenti}); err != nil {
			Log.Error("ComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(all{Genitori, Studenti, Docenti}); err != nil {
			Log.Error("ComunicatiHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(all{Genitori, Studenti, Docenti}); err != nil {
			Log.Error("ComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
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

	w.WriteHeader(http.StatusOK)
	switch RequestMime(r.Header) {

	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(GetByName("genitori")[:localCount]); err != nil {
			Log.Error("GenitoriComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(GetByName("genitori")[:localCount]); err != nil {
			Log.Error("GenitoriComunicatiHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(GetByName("genitori")[:localCount]); err != nil {
			Log.Error("GenitoriComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
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

	w.WriteHeader(http.StatusOK)
	switch RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(GetByName("studenti")[:localCount]); err != nil {
			Log.Error("StudentiComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(GetByName("studenti")[:localCount]); err != nil {
			Log.Error("StudentiComunicatiHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(GetByName("studenti")[:localCount]); err != nil {
			Log.Error("StudentiComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
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

	w.WriteHeader(http.StatusOK)
	switch RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(GetByName("docenti")[:localCount]); err != nil {
			Log.Error("DocentiComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(GetByName("docenti")[:localCount]); err != nil {
			Log.Error("DocentiComunicatiHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(GetByName("docenti")[:localCount]); err != nil {
			Log.Error("DocentiComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
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
	w.WriteHeader(http.StatusNotImplemented)
}
