package main

import (
	"flag"
	"net/http/fcgi"
)

var configPtr *string

func main() {
	configPtr = flag.String("config", "./config.toml", "Indirizzo del file di configurazione, in .toml o .json")
	flag.Parse()
	err := LoadPrefs(*configPtr)
	if err != nil {
		panic(err)
	}

	InitLogger()
	initServer()

	if GetConfig().Conn.FastCGI {
		router := NewRouter()
		Log.Fatal(fcgi.Serve(nil, router))
	} else {
		Log.Fatal(NewServer().ListenAndServe())
	}
}

func initServer() {
	var (
		GenitoriWatcher = FileWatcher{GetConfig().Dirs.Genitori, Genitori, func() {
			LoadComunicati(TipoGenitori)
		}}
		StudentiWatcher = FileWatcher{GetConfig().Dirs.Studenti, Studenti, func() {
			LoadComunicati(TipoStudenti)
		}}
		DocentiWatcher = FileWatcher{GetConfig().Dirs.Docenti, Docenti, func() {
			LoadComunicati(TipoDocenti)
		}}
	)
	Log.Info("---------- DaVinci API ----------")
	Log.Info("Avvio server...")
	Log.Info("Carico preferenze...")
	LoadComunicati(TipoGenitori)
	go GenitoriWatcher.Watch()
	LoadComunicati(TipoStudenti)
	go StudentiWatcher.Watch()
	LoadComunicati(TipoDocenti)
	go DocentiWatcher.Watch()

	Log.Info("Avvio completato.")
}
