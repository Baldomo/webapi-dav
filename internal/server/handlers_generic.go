package server

import (
	"encoding/json"
	"github.com/Baldomo/webapi-dav/internal/config"
	. "github.com/Baldomo/webapi-dav/internal/log"
	"github.com/Baldomo/webapi-dav/internal/utils"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

var (
	indexHtml = ""
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if indexHtml == "" {
		absPath, _ := filepath.Abs(filepath.Join(config.GetConfig().Dirs.HTML, "index.html"))
		raw, _ := ioutil.ReadFile(absPath)
		indexHtml = string(raw)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	temp, _ := template.New("index").Parse(indexHtml)
	temp.Execute(w, utils.TemplateData())
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
			Log.Error("AboutHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	case "text/html":
		if err := utils.ShowGenericTemplate(w, aboutMessage); err != nil {
			Log.Error("AboutHandler: errore template html")
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
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
			Log.Error(err.Error())
		}
		break

	case "application/json":
		if err := json.NewEncoder(w).Encode(message); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
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
		break

	case "text/html":
		if err := utils.ShowGenericTemplate(w, message); err != nil {
			Log.Error("TeapotHandler: errore template html")
		} else {
			w.WriteHeader(http.StatusTeapot)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	var versionMessage = utils.APIMessage{Code: http.StatusOK, Info: "webapi-dav v" + utils.VersionNumber}

	defer r.Body.Close()
	switch utils.RequestMime(r.Header) {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(versionMessage); err != nil {
			Log.Error("VersionHandler: errore encoding json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	case "text/html":
		if err := utils.ShowGenericTemplate(w, versionMessage); err != nil {
			Log.Error("VersionHandler: errore template html")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		break

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

func RefreshHTML() {
	Log.Info("Ricaricamento pagine web...")
	indexHtml = ""
	absPath, _ := filepath.Abs(filepath.Join(config.GetConfig().Dirs.HTML, "index.html"))
	raw, _ := ioutil.ReadFile(absPath)
	indexHtml = string(raw)
}
