package main

type Operations struct {
	Ops []*Operation
}

type Operation struct {
	URI   string
	Desc  string
	Title string
}

var (
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

	// TODO
	codes = map[int]string{
		404: "Elemento non trovato",
		400: "",
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
