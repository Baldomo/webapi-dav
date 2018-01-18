// +build windows

package main

import (
	"gopkg.in/hlandau/service.v2"
	"gopkg.in/hlandau/easyconfig.v1"
)

func main() {
	easyconfig.ParseFatal(nil, nil)

	service.Main(&service.Info{
		Title: "WebAPI Dav",
		Name: "webapi-dav",
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

			go StartServers()

			smgr.SetStarted()
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