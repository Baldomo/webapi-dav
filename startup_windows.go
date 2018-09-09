// +build windows

package main

import (
	"flag"
	"fmt"
	"github.com/Baldomo/webapi-dav/agenda"
	com "github.com/Baldomo/webapi-dav/comunicati"
	"github.com/Baldomo/webapi-dav/config"
	. "github.com/Baldomo/webapi-dav/log"
	"github.com/Baldomo/webapi-dav/orario"
	"github.com/Baldomo/webapi-dav/server"
	"github.com/Baldomo/webapi-dav/utils"
	"github.com/nightlyone/lockfile"
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
		GenitoriWatcher = FileWatcher{config.GetConfig().Dirs.Genitori, func() {
			com.LoadComunicati(com.TipoGenitori)
		}, true}
		StudentiWatcher = FileWatcher{config.GetConfig().Dirs.Studenti, func() {
			com.LoadComunicati(com.TipoStudenti)
		}, true}
		DocentiWatcher = FileWatcher{config.GetConfig().Dirs.Docenti, func() {
			com.LoadComunicati(com.TipoDocenti)
		}, true}
		OrarioWatcher = FileWatcher{config.GetConfig().Dirs.Orario, func() {
			orario.LoadOrario(config.GetConfig().Dirs.Orario)
		}, false}
		HTMLWatcher = WebContentWatcher{config.GetConfig().Dirs.HTML, func() {
			server.RefreshHTML()
		}}
		PrefWatcher = ConfigWatcher{config.GetConfigPath(), func() {
			config.ReloadPrefs()
		}}
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
