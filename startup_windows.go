// +build windows

package main

import (
	"flag"
	"fmt"
	"github.com/nightlyone/lockfile"
	"net/rpc"
	"os"
	"path/filepath"
)

const (
	pidfile = "webapi.pid"
)

var SH = new(ServerHandler)

func main() {
	configPtr := flag.String("config", "./config.toml", "Indirizzo del file di configurazione, in .toml o .json")
	versionPtr := flag.Bool("version", false, "Mostra la versione attuale del programma")
	stopPtr := flag.Bool("stop", false, "Termina l'esecuzione del programma")
	flag.Parse()

	if *stopPtr {
		client, err := rpc.Dial("tcp", ":2202")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		err = client.Call("ServerHandler.Shutdown", &struct{}{}, &struct{}{})
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	}

	if *versionPtr {
		fmt.Println("DaVinci API v" + VersionNumber)
		fmt.Println("Leonardo Baldin, " + VersionDate)
		os.Exit(0)
	}

	err := LoadPrefs(*configPtr)
	if err != nil {
		panic(err)
	}

	lockProcess()

	InitLogger(initServer)

	SH.Start()
}

func initServer() {
	var (
		GenitoriWatcher = FileWatcher{GetConfig().Dirs.Genitori, Genitori, func() {
			LoadComunicati(TipoGenitori)
		}, true}
		StudentiWatcher = FileWatcher{GetConfig().Dirs.Studenti, Studenti, func() {
			LoadComunicati(TipoStudenti)
		}, true}
		DocentiWatcher = FileWatcher{GetConfig().Dirs.Docenti, Docenti, func() {
			LoadComunicati(TipoDocenti)
		}, true}
		HTMLWatcher = WebContentWatcher{GetConfig().Dirs.HTML, func() {
			RefreshHTML()
		}}
		PrefWatcher = ConfigWatcher{GetConfigPath(), func() {
			ReloadPrefs()
		}}
	)
	Log.Info("---------- DaVinci API ----------")
	Log.Info("Avvio server...")
	Log.Info("Caricamento contenuti web...")
	go HTMLWatcher.Watch()

	Log.Info("Caricamento comunicati...")
	LoadComunicati(TipoGenitori)
	go GenitoriWatcher.Watch()
	LoadComunicati(TipoStudenti)
	go StudentiWatcher.Watch()
	LoadComunicati(TipoDocenti)
	go DocentiWatcher.Watch()

	Log.Info("Caricamento config...")
	go PrefWatcher.Watch()
	Log.Info("Avvio completato.")
	Log.Info("---------------------------------")
}

func lockProcess() {
	ex, _ := os.Executable()
	lock, err := lockfile.New(filepath.Join(filepath.Dir(ex), pidfile))
	if err != nil {
		panic(err)
	}

	err = lock.TryLock()
	if err != nil {
		os.Exit(1)
	}
}
