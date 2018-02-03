package server

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	// Generiche
	Route{
		"Index",
		"GET",
		"/api",
		IndexHandler,
	},
	Route{
		"VersionNumber",
		"GET",
		"/api/version",
		VersionHandler,
	},
	Route{
		"Informazioni",
		"GET",
		"/api/about",
		AboutHandler,
	},
	Route{
		"Teapot",
		"GET",
		"/api/teapot",
		TeapotHandler,
	},

	// Comunicati
	Route{
		"Comunicati_List",
		"GET",
		"/api/comunicati",
		ComunicatiHandler,
	},
	Route{
		"Comunicati_List_Genitori",
		"GET",
		"/api/comunicati/genitori",
		GenitoriComunicatiHandler,
	},
	Route{
		"Comunicati_List_Genitori",
		"GET",
		"/api/comunicati/genitori/{count:[0-9]+}",
		GenitoriComunicatiHandler,
	},
	Route{
		"Comunicati_List_Studenti",
		"GET",
		"/api/comunicati/studenti",
		StudentiComunicatiHandler,
	},
	Route{
		"Comunicati_List_Studenti",
		"GET",
		"/api/comunicati/studenti/{count:[0-9]+}",
		StudentiComunicatiHandler,
	},
	Route{
		"Comunicati_List_Docenti",
		"GET",
		"/api/comunicati/docenti",
		DocentiComunicatiHandler,
	},
	Route{
		"Comunicati_List_Docenti",
		"GET",
		"/api/comunicati/docenti/{count:[0-9]+}",
		DocentiComunicatiHandler,
	},

	// Utilità
	Route{
		"Docenti_List",
		"GET",
		"/api/docenti",
		DocentiHandler,
	},

	// Orario
	Route{
		"Orario_Table",
		"GET",
		"/api/orario",
		OrarioHandler,
	},
	/*Route{
		"Orario_Table_Giorno",
		"GET",
		`/api/orario/{giorno:(?:(?:(luned)|(marted)|(mercoled)|(gioved)|(venerd))([i|ì]))|(sabato)}`,
		OrarioGiornoHandler,
	},*/
	Route{
		"Orario_Table_Classe",
		"GET",
		"/api/orario/classe/{classe:[1-5][a-zA-Z]}",
		OrarioClasseHandler,
	},
	Route{
		"Orario_Table_Docente",
		"GET",
		"/api/orario/docente/{cognome:[a-zA-Z]+}",
		OrarioDocenteHandler,
	},

	// Progetti
	Route{
		"Progetti_List",
		"GET",
		"/api/progetti",
		ProgettiHandler,
	},
}
