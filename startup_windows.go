// +build windows

package main

import (
	"flag"
	"fmt"
	"github.com/nightlyone/lockfile"
	"os"
	"path/filepath"
	"time"
)

const (
	pidfile = "webapi.pid"
)

func main() {
	configPtr := flag.String("config", "./config.toml", "Indirizzo del file di configurazione, in .toml o .json")
	versionPtr := flag.Bool("version", false, "Mostra la versione attuale del programma")
	closePtr := flag.Bool("close", false, "Termina l'esecuzione del programma")
	flag.Parse()

	if *closePtr {
		ex, _ := os.Executable()
		lock, errLock := lockfile.New(filepath.Join(filepath.Dir(ex), pidfile))
		if errLock != nil {
			panic(errLock)
		}

		process, errProc := lock.GetOwner()
		if errProc != nil {
			// Owner dead ErrDeadOwner
			os.Exit(1)
		} else {
			process.Signal(os.Interrupt)
			time.Sleep(3 * time.Second)
			process.Kill()
			os.Exit(0)
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

	StartServers()
}

func initServer() {
	var (
		GenitoriWatcher = FileWatcher{GetConfig().Dirs.Genitori, Genitori, func() {
			LoadComunicati(TipoGenitori)
		}, true, ComunicatiWatcher}
		StudentiWatcher = FileWatcher{GetConfig().Dirs.Studenti, Studenti, func() {
			LoadComunicati(TipoStudenti)
		}, true, ComunicatiWatcher}
		DocentiWatcher = FileWatcher{GetConfig().Dirs.Docenti, Docenti, func() {
			LoadComunicati(TipoDocenti)
		}, true, ComunicatiWatcher}
		HTMLWatcher = WebContentWatcher{GetConfig().Dirs.HTML, func() {
			RefreshHTML()
		}}
		PrefWatcher = ConfigWatcher{GetConfigPath(), func() {
			LoadPrefs(GetConfigPath())
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
