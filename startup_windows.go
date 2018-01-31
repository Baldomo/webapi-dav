// +build windows

package main

import (
	"flag"
	"fmt"
	"github.com/nightlyone/lockfile"
	. "leonardobaldin/webapi-dav/config"
	"leonardobaldin/webapi-dav/server"
	"leonardobaldin/webapi-dav/sezioni"
	"leonardobaldin/webapi-dav/utils"
	"net/rpc"
	"os"
	"path/filepath"
)

const (
	pidfile = "webapi.pid"
)

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

		err = client.Call("serverHandler.Shutdown", &struct{}{}, &struct{}{})
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	}

	if *versionPtr {
		fmt.Println("DaVinci API v" + utils.VersionNumber)
		fmt.Println("Leonardo Baldin, " + utils.VersionDate)
		os.Exit(0)
	}

	err := LoadPrefs(*configPtr)
	if err != nil {
		panic(err)
	}

	lockProcess()

	InitLogger(initServer)

	server.Handler.Start()
}

func initServer() {
	var (
		GenitoriWatcher = FileWatcher{GetConfig().Dirs.Genitori, sezioni.Genitori, func() {
			sezioni.LoadComunicati(sezioni.TipoGenitori)
		}, true, ComunicatiWatcher}
		StudentiWatcher = FileWatcher{GetConfig().Dirs.Studenti, sezioni.Studenti, func() {
			sezioni.LoadComunicati(sezioni.TipoStudenti)
		}, true, ComunicatiWatcher}
		DocentiWatcher = FileWatcher{GetConfig().Dirs.Docenti, sezioni.Docenti, func() {
			sezioni.LoadComunicati(sezioni.TipoDocenti)
		}, true, ComunicatiWatcher}
		HTMLWatcher = WebContentWatcher{GetConfig().Dirs.HTML, func() {
			server.RefreshHTML()
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
	sezioni.LoadComunicati(sezioni.TipoGenitori)
	go GenitoriWatcher.Watch()
	sezioni.LoadComunicati(sezioni.TipoStudenti)
	go StudentiWatcher.Watch()
	sezioni.LoadComunicati(sezioni.TipoDocenti)
	go DocentiWatcher.Watch()

	Log.Info("Caricamento orario...")
	sezioni.LoadOrario(GetConfig().Dirs.Orario)

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
