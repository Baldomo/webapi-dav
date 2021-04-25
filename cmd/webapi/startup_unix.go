//+build linux darwin

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	com "github.com/Baldomo/webapi-dav/pkg/comunicati"
	"github.com/Baldomo/webapi-dav/pkg/config"
	"github.com/Baldomo/webapi-dav/pkg/auth"
	"github.com/Baldomo/webapi-dav/pkg/log"
	"github.com/Baldomo/webapi-dav/pkg/orario"
	"github.com/Baldomo/webapi-dav/pkg/server"
	"github.com/Baldomo/webapi-dav/pkg/utils"
	"github.com/Baldomo/webapi-dav/pkg/watchers"
	"github.com/nightlyone/lockfile"
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
		ex, _ := os.Executable()
		lock, errLock := lockfile.New(filepath.Join(filepath.Dir(ex), pidfile))
		if errLock != nil {
			panic(errLock)
		}

		proc, errProc := lock.GetOwner()
		if errProc != nil {
			// Owner dead ErrDeadOwner
			os.Exit(1)
		} else {
			proc.Signal(os.Interrupt)
			//time.Sleep(3 * time.Second)
			//process.Kill()
			os.Exit(0)
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

    server.Init()

	lockProcess()

	err = auth.InitializeSigning()
	if err != nil {
		panic(err)
	}

	log.InitLogger()
	initServer()

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
	log.Info("---------- DaVinci API ----------")
	log.Info("Avvio server...")
	log.Info("Caricamento contenuti web...")
	go HTMLWatcher.Watch()

	log.Info("Caricamento comunicati...")
	com.LoadComunicati(com.TipoGenitori)
	go GenitoriWatcher.Watch()
	com.LoadComunicati(com.TipoStudenti)
	go StudentiWatcher.Watch()
	com.LoadComunicati(com.TipoDocenti)
	go DocentiWatcher.Watch()

	log.Info("Caricamento orario...")
	orario.LoadOrario(config.GetConfig().Dirs.Orario)
	go OrarioWatcher.Watch()

	log.Info("Caricamento config...")
	go PrefWatcher.Watch()
	log.Info("Avvio completato.")
	log.Info("---------------------------------")
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
