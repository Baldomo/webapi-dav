package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Baldomo/webapi-dav/pkg/agenda"
	com "github.com/Baldomo/webapi-dav/pkg/comunicati"
	"github.com/Baldomo/webapi-dav/pkg/log"
	"github.com/Baldomo/webapi-dav/pkg/orario"
	"github.com/Baldomo/webapi-dav/pkg/utils"
	"github.com/gorilla/mux"
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

	defer r.Body.Close()
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(all); err != nil {
			log.Error("ComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(all); err != nil {
			log.Error("ComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func GenitoriComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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
			log.Error("GenitoriComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(com.GetByName("genitori")[:localCount]); err != nil {
			log.Error("GenitoriComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func StudentiComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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
			log.Error("StudentiComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(com.GetByName("studenti")[:localCount]); err != nil {
			log.Error("StudentiComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func DocentiComunicatiHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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
			log.Error("DocentiComunicatiHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(com.GetByName("docenti")[:localCount]); err != nil {
			log.Error("DocentiComunicatiHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

// Utilità

func DocentiHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetAllDocenti()); err != nil {
			log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetAllDocenti()); err != nil {
			log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func ClassiHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetAllClassi()); err != nil {
			log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetAllClassi()); err != nil {
			log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

// Orario

func OrarioHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetOrario()); err != nil {
			log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetOrario()); err != nil {
			log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func OrarioClasseHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	classe := mux.Vars(r)["classe"]
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetByClasse(classe)); err != nil {
			log.Error("OrarioClasseHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetByClasse(classe)); err != nil {
			log.Error("OrarioClasseHandler: errore encoding html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}

}

func OrarioDocenteHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var doc orario.Docente
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := doc.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetByDoc(doc)); err != nil {
			log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	case "text/html":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(orario.GetByDoc(doc)); err != nil {
			log.Error("OrarioHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

// Agenda

func AgendaHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var es agenda.EventStream
	if err := json.NewDecoder(r.Body).Decode(&es); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if utils.RequestMime(r.Header) == "application/json" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		events, err := es.Close()
		if err != nil {
			log.Error("AgendaHandler: errore EventStream.Close()")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(events); err != nil {
			log.Error("AgendaHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
