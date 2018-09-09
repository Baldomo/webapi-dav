package server

import (
	"encoding/json"
	"github.com/Baldomo/webapi-dav/agenda"
	com "github.com/Baldomo/webapi-dav/comunicati"
	. "github.com/Baldomo/webapi-dav/log"
	"github.com/Baldomo/webapi-dav/orario"
	"github.com/Baldomo/webapi-dav/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Comunicati

func ComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	all := struct {
		Genitori com.Comunicati `json:"genitori"`
		Studenti com.Comunicati `json:"studenti"`
		Docenti  com.Comunicati `json:"docenti"`
	}{
		com.Genitori,
		com.Studenti,
		com.Docenti,
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
		break

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(all); err != nil {
			Log.Error("ComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
	return
}

func GenitoriComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	var localCount int
	vars := mux.Vars(r)
	count, countPresent := vars["count"]
	if !countPresent {
		localCount = com.GetLenByName("genitori")
	} else {
		localCount, _ = strconv.Atoi(count)
		if localCount > com.GetLenByName("genitori") {
			localCount = com.GetLenByName("genitori")
		}
	}

	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(com.GetByName("genitori")[:localCount]); err != nil {
			Log.Error("GenitoriComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(com.GetByName("genitori")[:localCount]); err != nil {
			Log.Error("GenitoriComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

func StudentiComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	var localCount int
	vars := mux.Vars(r)
	count, countPresent := vars["count"]
	if !countPresent {
		localCount = com.GetLenByName("studenti")
	} else {
		localCount, _ = strconv.Atoi(count)
		if localCount > com.GetLenByName("studenti") {
			localCount = com.GetLenByName("studenti")
		}
	}

	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(com.GetByName("studenti")[:localCount]); err != nil {
			Log.Error("StudentiComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(com.GetByName("studenti")[:localCount]); err != nil {
			Log.Error("StudentiComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

func DocentiComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	var localCount int
	vars := mux.Vars(r)
	count, countPresent := vars["count"]
	if !countPresent {
		localCount = com.GetLenByName("docenti")
	} else {
		localCount, _ = strconv.Atoi(count)
		if localCount > com.GetLenByName("docenti") {
			localCount = com.GetLenByName("docenti")
		}
	}

	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(com.GetByName("docenti")[:localCount]); err != nil {
			Log.Error("DocentiComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(com.GetByName("docenti")[:localCount]); err != nil {
			Log.Error("DocentiComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

// Utilità

func DocentiHandler(w http.ResponseWriter, r *http.Request) {
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetAllDocenti()); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetAllDocenti()); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

func ClassiHandler(w http.ResponseWriter, r *http.Request) {
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetAllClassi()); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetAllClassi()); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

// Orario

func OrarioHandler(w http.ResponseWriter, r *http.Request) {
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetOrario()); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetOrario()); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

func OrarioClasseHandler(w http.ResponseWriter, r *http.Request) {
	classe, _ := mux.Vars(r)["classe"]

	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetByClasse(classe)); err != nil {
			Log.Error("OrarioClasseHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetByClasse(classe)); err != nil {
			Log.Error("OrarioClasseHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}

}

func OrarioDocenteHandler(w http.ResponseWriter, r *http.Request) {
	var doc orario.Docente
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&doc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetByDoc(doc)); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetByDoc(doc)); err != nil {
			Log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

// Agenda

func AgendaHandler(w http.ResponseWriter, r *http.Request) {
	var es agenda.EventStream
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&es)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if utils.RequestMime(r.Header) == "application/json" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(es.Close()); err != nil {
			Log.Error("AgendaHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	} else {
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

// Progetti

func ProgettiHandler(w http.ResponseWriter, r *http.Request) {
	//TODO prende da file toml/json tutto il necessario, slice di oggetti
	type progetto struct {
		Nome         string `json:"nome" toml:"nome"`
		PathImmagine string `json:"immagine" toml:"immagine"`
		Descr        string `json:"descrizione" toml:"descrizione"`
	}
	w.WriteHeader(http.StatusNotFound)
}
