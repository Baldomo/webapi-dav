package main

import (
	"html/template"
	"net/http"
)

type Operation struct {
	URI   string
	Desc  string
	Title string
}

type APIMessage struct {
	Code uint   `json:"codice" xml:"codice"`
	Info string `json:"info" xml:"info"`
}

const (
	Version = "0.1.0"
)

var (
	genericHTML = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Code}} - DaVinci API</title>
		<meta charset="utf-8">
		<meta name="og:description" content="{{.Code}} - {{.Info}}">
		<meta name="twitter:card" content="summary" />
		<link href="//netdna.bootstrapcdn.com/twitter-bootstrap/2.3.2/css/bootstrap-combined.min.css" rel="stylesheet">
	</head>
	<body style="">
	<div class="container">
		<div class="hero-unit">
			<h1>{{.Code}} - {{.Info}}</h1>
		</div>
	</div>
	</body>
	</html>`

	ops = map[string]*Operation{
		"about":                   {"/api/about", "Restituirà informazioni generali sulla API", "/about"},
		"comunicati":              {"/api/comunicati", "Restituirà la lista completa di comunicati", "/comunicati"},
		"comunicati-docenti":      {"/api/comunicati/docenti", "Restituirà la lista dei comunicati per i docenti", "/comunicati/docenti"},
		"comunicati-docenti-arg":  {"/api/comunicati/docenti/5", "Restituirà gli ultimi  n comunicati per i docenti", "/comunicati/docenti/{n: numero}"},
		"comunicati-genitori":     {"/api/comunicati/genitori", "Restituirà la lista dei comunicati per i genitori", "/comunicati/genitori"},
		"comunicati-genitori-arg": {"/api/comunicati/genitori/5", "Restituirà gli ultimi n comunicati per i genitori", "/comunicati/genitori/{n: numero}"},
		"comunicati-studenti":     {"/api/comunicati/studenti", "Restituirà la lista dei comunicati per gli studenti", "/comunicati/studenti"},
		"comunicati-studenti-arg": {"/api/comunicati/studenti/5", "Restituirà gli ultimi n comunicati per gli studenti", "/comunicati/studenti/{n: numero}"},
		"version":                 {"/api/version", "Restituirà la versione dell'API in uso", "/version"},
	}
)

func GetOp(nome string) (*Operation, error) {
	if val, ok := ops[nome]; ok {
		return val, nil
	} else {
		return nil, Error("GetOp: ", "Valore non presente %s", nome)
	}
}

func GetMapOps() *map[string]*Operation {
	return &ops
}

func ShowGenericTemplate(w http.ResponseWriter, args interface{}) error {
	temp, err := template.New("generic").Parse(genericHTML)
	if err != nil {
		return err
	}
	temp.Execute(w, args)
	return nil
}