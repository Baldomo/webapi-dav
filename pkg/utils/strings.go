package utils

import (
	"html/template"
	"net/http"
)

// Modello di informazione di un endpoint
type Operation struct {
	// Metodo accettato (ad es. "POST")
	Method string

	// URL dell'endpoint (ad es. "/api/orario")
	URI string

	// Descrizione testuale della funzionalità
	// dell'endpoint
	Desc string

	// Nome simbolico dell'endpoint
	Title string
}

// Modello di un messaggio generico di ritorno per una richiesta agli
// endpoint di utility (ad es. /api/version)
type APIMessage struct {
	// Codice HTTP (ad es. 200 o 404)
	Code uint `json:"codice"`

	// Stringa con informazioni aggiuntive
	Info string `json:"info"`
}

// Rappresentazione interna dei dati passati alla template di index.html
// visualizzata nel browser per le richieste all'URL di base (/api)
type templateData struct {
	Version string
	Ops     map[string]*Operation
}

const (
	// Versione corrente di webapi-dav
	VersionNumber = "0.7.1"

	// Data di release della versione corrente
	VersionDate = "08/09/2019"
)

var (
	// Template generica mostrata per le richieste da browser a endpoint di utility
	// (ad es. /api/version)
	genericHTML = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Code}} - DaVinci API</title>
		<meta charset="utf-8">
		<meta name="og:description" content="{{.Code}} - {{.Info}}">
		<meta name="twitter:card" content="summary" />
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" 
			integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
	</head>
	<body style="">
	<div class="container">
		<div class="jumbotron">
			<h1>{{.Code}} - {{.Info}}</h1>
		</div>
	</div>
	</body>
	</html>`

	ops = map[string]*Operation{
		"about":                   {"GET", "/api/about", "Restituirà informazioni generali sulla API", "/about"},
		"agenda":                  {"POST", "/api/agenda", "Restituisce eventi dell'agenda filtrati dal contenuto della richiesta POST JSON", "/agenda"},
		"classi":                  {"GET", "/api/classi", "Restituirà la lista di tutte le classi del liceo", "/classi"},
		"comunicati":              {"GET", "/api/comunicati", "Restituirà la lista completa di comunicati", "/comunicati"},
		"comunicati-docenti":      {"GET", "/api/comunicati/docenti", "Restituirà la lista dei comunicati per i docenti", "/comunicati/docenti"},
		"comunicati-docenti-arg":  {"GET", "/api/comunicati/docenti/5", "Restituirà gli ultimi n comunicati per i docenti", "/comunicati/docenti/{n: numero}"},
		"comunicati-genitori":     {"GET", "/api/comunicati/genitori", "Restituirà la lista dei comunicati per i genitori", "/comunicati/genitori"},
		"comunicati-genitori-arg": {"GET", "/api/comunicati/genitori/5", "Restituirà gli ultimi n comunicati per i genitori", "/comunicati/genitori/{n: numero}"},
		"comunicati-studenti":     {"GET", "/api/comunicati/studenti", "Restituirà la lista dei comunicati per gli studenti", "/comunicati/studenti"},
		"comunicati-studenti-arg": {"GET", "/api/comunicati/studenti/5", "Restituirà gli ultimi n comunicati per gli studenti", "/comunicati/studenti/{n: numero}"},
		"docenti":                 {"GET", "/api/docenti", "Restituirà la lista dei docenti del liceo", "/docenti"},
		"orario":                  {"GET", "/api/orario", "Restituirà l'orario completo di tutte le classi (PESANTE)", "/orario"},
		"orario-classe":           {"GET", "/api/orario/classe/4b", "Restituirà l'orario della classe specificata", "/orario/classe/{c: classe}"},
		"orario-docente":          {"POST", "/api/orario/docente", "Restituisce l'orario del docente con nome e cognome nella richiesta POST in JSON", "/orario/docente..."},
		"teapot":                  {"GET", "/api/teapot", "Restituirà codice HTTP 418. Utile solamente a capire se la API è online e funzionante", "/teapot"},
		"version":                 {"GET", "/api/version", "Restituirà la versione dell'API in uso", "/version"},
	}
)

// Costruttore per templateData
func TemplateData() templateData {
	return templateData{
		Ops:     ops,
		Version: VersionNumber,
	}
}

// Restituisce una Operation memorizzata dato il nome associato
func GetOp(nome string) *Operation {
	if val, ok := ops[nome]; ok {
		return val
	} else {
		return nil
	}
}

// Restituisce la map interna contenente tutti i dati per index.html
func GetMapOps() *map[string]*Operation {
	return &ops
}

// Dato un http.ReponseWriter, scrive in uscita la template generica
// con gli argomenti dati (di tipo APIMessage)
func ShowGenericTemplate(w http.ResponseWriter, args APIMessage) error {
	temp, err := template.New("generic").Parse(genericHTML)
	if err != nil {
		return err
	}
	return temp.Execute(w, args)
}
