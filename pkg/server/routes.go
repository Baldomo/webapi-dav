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

	// Richiede autorizzazione
	RequiresAuthorization bool

	// Funzione handler associata
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	// Generiche
	{
		"Index",
		"GET",
		"/api",
		false,
		IndexHandler,
	},
	{
		"VersionNumber",
		"GET",
		"/api/version",
		false,
		VersionHandler,
	},
	{
		"Informazioni",
		"GET",
		"/api/about",
		false,
		AboutHandler,
	},
	{
		"Teapot",
		"GET",
		"/api/teapot",
		false,
		TeapotHandler,
	},
	{
		"Swagger",
		"GET",
		"/api/openapi.yaml",
		false,
		OpenapiHandler,
	},

	// Comunicati
	{
		"Comunicati_List",
		"GET",
		"/api/comunicati",
		true,
		ComunicatiHandler,
	},
	{
		"Comunicati_List_Genitori",
		"GET",
		"/api/comunicati/genitori",
		true,
		GenitoriComunicatiHandler,
	},
	{
		"Comunicati_List_Genitori",
		"GET",
		"/api/comunicati/genitori/{count:[0-9]+}",
		true,
		GenitoriComunicatiHandler,
	},
	{
		"Comunicati_List_Studenti",
		"GET",
		"/api/comunicati/studenti",
		true,
		StudentiComunicatiHandler,
	},
	{
		"Comunicati_List_Studenti",
		"GET",
		"/api/comunicati/studenti/{count:[0-9]+}",
		true,
		StudentiComunicatiHandler,
	},
	{
		"Comunicati_List_Docenti",
		"GET",
		"/api/comunicati/docenti",
		true,
		DocentiComunicatiHandler,
	},
	{
		"Comunicati_List_Docenti",
		"GET",
		"/api/comunicati/docenti/{count:[0-9]+}",
		true,
		DocentiComunicatiHandler,
	},

	// Utilità
	{
		"Docenti_List",
		"GET",
		"/api/docenti",
		true,
		DocentiHandler,
	},
	{
		"Classi_List",
		"GET",
		"/api/classi",
		true,
		ClassiHandler,
	},
	{
		"Pdf",
		"GET",
		"/sitoLiceo/images/comunicati/comunicati-docenti/{filename}",
		true,
		PdfHandler,
	},
	{
		"Pdf",
		"GET",
		"/sitoLiceo/images/comunicati/comunicati-genitori/{filename}",
		true,
		PdfHandler,
	},
	{
		"Pdf",
		"GET",
		"/sitoLiceo/images/comunicati/comunicati-studenti/{filename}",
		true,
		PdfHandler,
	},

	// Orario
	{
		"Orario_Table",
		"GET",
		"/api/orario",
		true,
		OrarioHandler,
	},
	{
		"Orario_Table_Classe",
		"GET",
		"/api/orario/classe/{classe:[1-5][a-zA-Z]}",
		true,
		OrarioClasseHandler,
	},
	{
		"Orario_Table_Docente",
		"POST",
		"/api/orario/docente",
		true,
		OrarioDocenteHandler,
	},

	// Agenda
	{
		"Agenda",
		"POST",
		"/api/agenda",
		false,
		AgendaHandler,
	},
}
