package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	configPtr := flag.String("config", "./config.toml", "Indirizzo del file di configurazione, in .toml o .json")
	versionPtr := flag.Bool("version", false, "Mostra la versione attuale del programma")
	flag.Parse()

	if *versionPtr {
		fmt.Println("DaVinci API v" + VersionNumber)
		fmt.Println("Leonardo Baldin, " + VersionDate)
		os.Exit(0)
	}

	err := LoadPrefs(*configPtr)
	if err != nil {
		panic(err)
	}

	InitLogger(initServer)

	StartServers()
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
	Log.Info("---------------------------------")
}
