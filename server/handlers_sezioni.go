package server

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	. "leonardobaldin/webapi-dav/log"
	"leonardobaldin/webapi-dav/sezioni"
	"leonardobaldin/webapi-dav/utils"
	"net/http"
	"strconv"
)

// Comunicati

func ComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	all := struct {
		Genitori sezioni.Comunicati `json:"genitori" xml:"genitori"`
		Studenti sezioni.Comunicati `json:"studenti" xml:"studenti"`
		Docenti  sezioni.Comunicati `json:"docenti" xml:"docenti"`
	}{
		sezioni.Genitori,
		sezioni.Studenti,
		sezioni.Docenti,
	}

	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(all); err != nil {
			Log.Error("ComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(all); err != nil {
			Log.Error("ComunicatiHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(all); err != nil {
			Log.Error("ComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
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
		localCount = sezioni.GetLenByName("genitori")
	} else {
		localCount, _ = strconv.Atoi(count)
		if localCount > sezioni.GetLenByName("genitori") {
			localCount = sezioni.GetLenByName("genitori")
		}
	}

	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByName("genitori")[:localCount]); err != nil {
			Log.Error("GenitoriComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(sezioni.GetByName("genitori")[:localCount]); err != nil {
			Log.Error("GenitoriComunicatiHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByName("genitori")[:localCount]); err != nil {
			Log.Error("GenitoriComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
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
		localCount = sezioni.GetLenByName("studenti")
	} else {
		localCount, _ = strconv.Atoi(count)
		if localCount > sezioni.GetLenByName("studenti") {
			localCount = sezioni.GetLenByName("studenti")
		}
	}

	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByName("studenti")[:localCount]); err != nil {
			Log.Error("StudentiComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(sezioni.GetByName("studenti")[:localCount]); err != nil {
			Log.Error("StudentiComunicatiHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByName("studenti")[:localCount]); err != nil {
			Log.Error("StudentiComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
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
		localCount = sezioni.GetLenByName("docenti")
	} else {
		localCount, _ = strconv.Atoi(count)
		if localCount > sezioni.GetLenByName("docenti") {
			localCount = sezioni.GetLenByName("docenti")
		}
	}

	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByName("docenti")[:localCount]); err != nil {
			Log.Error("DocentiComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(sezioni.GetByName("docenti")[:localCount]); err != nil {
			Log.Error("DocentiComunicatiHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByName("docenti")[:localCount]); err != nil {
			Log.Error("DocentiComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Utilità
func DocentiHandler(w http.ResponseWriter, r *http.Request) {
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetAllDocenti()); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(sezioni.GetAllDocenti()); err != nil {
			Log.Error("OrarioHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetAllDocenti()); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Orario

func OrarioHandler(w http.ResponseWriter, r *http.Request) {
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetOrario()); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(sezioni.GetOrario()); err != nil {
			Log.Error("OrarioHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetOrario()); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func OrarioGiornoHandler(w http.ResponseWriter, r *http.Request) {
	giorno, _ := mux.Vars(r)["giorno"]
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByGiorno(giorno)); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(sezioni.GetByGiorno(giorno)); err != nil {
			Log.Error("OrarioHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByGiorno(giorno)); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func OrarioClasseHandler(w http.ResponseWriter, r *http.Request) {
	classe, _ := mux.Vars(r)["classe"]

	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByClasse(classe)); err != nil {
			Log.Error("OrarioClasseHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(sezioni.GetByClasse(classe)); err != nil {
			Log.Error("OrarioClasseHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByClasse(classe)); err != nil {
			Log.Error("OrarioClasseHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}

func OrarioDocenteHandler(w http.ResponseWriter, r *http.Request) {
	cogn, _ := mux.Vars(r)["cognome"]
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByDocCogn(cogn)); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if err := xml.NewEncoder(w).Encode(sezioni.GetByDocCogn(cogn)); err != nil {
			Log.Error("OrarioHandler: errore encoding xml")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(sezioni.GetByDocCogn(cogn)); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Progetti

func ProgettiHandler(w http.ResponseWriter, r *http.Request) {
	//TODO prende da file toml/json tutto il necessario, ARRAY di oggetti
	type progetto struct {
		Nome         string `json:"nome" toml:"nome"`
		PathImmagine string `json:"immagine" toml:"immagine"`
		Descr        string `json:"descrizione" toml:"descrizione"`
	}
	w.WriteHeader(http.StatusNotFound)
}
