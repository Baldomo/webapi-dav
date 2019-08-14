// +build windows

package main

import (
	"flag"
	"fmt"
	"github.com/Baldomo/webapi-dav/internal/agenda"
	com "github.com/Baldomo/webapi-dav/internal/comunicati"
	"github.com/Baldomo/webapi-dav/internal/config"
	. "github.com/Baldomo/webapi-dav/internal/log"
	"github.com/Baldomo/webapi-dav/internal/orario"
	"github.com/Baldomo/webapi-dav/internal/server"
	"github.com/Baldomo/webapi-dav/internal/utils"
	"github.com/Baldomo/webapi-dav/internal/watchers"
	"github.com/nightlyone/lockfile"
	"net/rpc"
	"os"
	"path/filepath"
)

const (
	pidfile = "webapi.pid"
)

func start() {
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

	err := config.LoadPrefs(*configPtr)
	if err != nil {
		panic(err)
	}

	lockProcess()

	InitLogger(initServer)

	server.Start()
}

func initServer() {
	var (
		GenitoriWatcher = watchers.FileWatcher{
			Path: config.GetConfig().Dirs.Genitori,
			OnEvent: func() {
				com.LoadComunicati(com.TipoGenitori)
			},
			Notify: true,
		}
		StudentiWatcher = watchers.FileWatcher{
			Path: config.GetConfig().Dirs.Studenti,
			OnEvent: func() {
				com.LoadComunicati(com.TipoStudenti)
			},
			Notify: true,
		}
		DocentiWatcher = watchers.FileWatcher{
			Path: config.GetConfig().Dirs.Docenti,
			OnEvent: func() {
				com.LoadComunicati(com.TipoDocenti)
			},
		Notify: true}
		OrarioWatcher = watchers.FileWatcher{
			Path: config.GetConfig().Dirs.Orario,
			OnEvent: func() {
				orario.LoadOrario(config.GetConfig().Dirs.Orario)
			},
		}
		HTMLWatcher = watchers.WebContentWatcher{
			Path: config.GetConfig().Dirs.HTML,
			OnEvent: func() {
				server.RefreshHTML()
			},
		}
		PrefWatcher = watchers.ConfigWatcher{
			Path: config.GetConfigPath(),
			OnEvent: func() {
				config.ReloadPrefs()
			},
		}
	)
	Log.Info("---------- DaVinci API ----------")
	Log.Info("Avvio server...")
	Log.Info("Caricamento contenuti web...")
	go HTMLWatcher.Watch()

	Log.Info("Caricamento comunicati...")
	com.LoadComunicati(com.TipoGenitori)
	go GenitoriWatcher.Watch()
	com.LoadComunicati(com.TipoStudenti)
	go StudentiWatcher.Watch()
	com.LoadComunicati(com.TipoDocenti)
	go DocentiWatcher.Watch()

	Log.Info("Caricamento orario...")
	orario.LoadOrario(config.GetConfig().Dirs.Orario)
	go OrarioWatcher.Watch()

	Log.Info("Caricamento config...")
	go PrefWatcher.Watch()

	Log.Info("Collegamento a database...")
	agenda.Fetch()

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
