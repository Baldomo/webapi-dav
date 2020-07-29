package server

import "net/http"

// Rappresentazione schematica di un endpoint REST
type Route struct {
	// Nome simbolico
	Name string

	// Metodo REST accettato
	Method string

	// Pattern testuale per l'endpoint (ad es. "/api/comunicati/genitori/{count:[0-9]+}")
	Pattern string

	// Funzione handler associata
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	// Generiche
	{
		"Index",
		"GET",
		"/api",
		IndexHandler,
	},
	{
		"VersionNumber",
		"GET",
		"/api/version",
		VersionHandler,
	},
	{
		"Informazioni",
		"GET",
		"/api/about",
		AboutHandler,
	},
	{
		"Teapot",
		"GET",
		"/api/teapot",
		TeapotHandler,
	},
	{
		"Swagger",
		"GET",
		"/api/openapi.yaml",
		OpenapiHandler,
	},

	// Comunicati
	{
		"Comunicati_List",
		"GET",
		"/api/comunicati",
		ComunicatiHandler,
	},
	{
		"Comunicati_List_Genitori",
		"GET",
		"/api/comunicati/genitori",
		GenitoriComunicatiHandler,
	},
	{
		"Comunicati_List_Genitori",
		"GET",
		"/api/comunicati/genitori/{count:[0-9]+}",
		GenitoriComunicatiHandler,
	},
	{
		"Comunicati_List_Studenti",
		"GET",
		"/api/comunicati/studenti",
		StudentiComunicatiHandler,
	},
	{
		"Comunicati_List_Studenti",
		"GET",
		"/api/comunicati/studenti/{count:[0-9]+}",
		StudentiComunicatiHandler,
	},
	{
		"Comunicati_List_Docenti",
		"GET",
		"/api/comunicati/docenti",
		DocentiComunicatiHandler,
	},
	{
		"Comunicati_List_Docenti",
		"GET",
		"/api/comunicati/docenti/{count:[0-9]+}",
		DocentiComunicatiHandler,
	},

	// Utilità
	{
		"Docenti_List",
		"GET",
		"/api/docenti",
		DocentiHandler,
	},
	{
		"Classi_List",
		"GET",
		"/api/classi",
		ClassiHandler,
	},

	// Orario
	{
		"Orario_Table",
		"GET",
		"/api/orario",
		OrarioHandler,
	},
	{
		"Orario_Table_Classe",
		"GET",
		"/api/orario/classe/{classe:[1-5][a-zA-Z]}",
		OrarioClasseHandler,
	},
	{
		"Orario_Table_Docente",
		"POST",
		"/api/orario/docente",
		OrarioDocenteHandler,
	},

	// Agenda
	{
		"Agenda",
		"POST",
		"/api/agenda",
		AgendaHandler,
	},
}
