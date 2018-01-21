// +build windows

package main

import (
	"gopkg.in/hlandau/easyconfig.v1"
	"gopkg.in/hlandau/service.v2"
	"flag"
	"fmt"
	"os"
)

func main() {
	easyconfig.ParseFatal(nil, nil)
	versionPtr := flag.Bool("version", false, "Mostra la versione attuale del programma")
	flag.Parse()

	if *versionPtr {
		fmt.Println("DaVinci API v" + VersionNumber)
		fmt.Println("Leonardo Baldin, " + VersionDate)
		os.Exit(0)
	}

	service.Main(&service.Info{
		Title:       "WebAPI Dav",
		Name:        "webapi-dav",
		Description: "Servizio per gestione web API",

		RunFunc: func(smgr service.Manager) error {

			err := smgr.DropPrivileges()
			if err != nil {
				return err
			}

			err = LoadPrefs("config.toml")
			if err != nil {
				return err
			}

			InitLogger(initServer)

			StartServers()

			smgr.SetStarted()
			smgr.SetStatus("webapi-dav in esecuzione")

			<-smgr.StopChan()

			return nil
		},
	})
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
