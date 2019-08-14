package watchers

import (
	com "github.com/Baldomo/webapi-dav/internal/comunicati"
	"github.com/Baldomo/webapi-dav/internal/config"
	. "github.com/Baldomo/webapi-dav/internal/log"
	"github.com/radovskyb/watcher"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type WatcherType uint64

type Watcher interface {
	Watch()
}

type FileWatcher struct {
	Path    string
	OnEvent func()
	Notify  bool
}

type WebContentWatcher struct {
	Path    string
	OnEvent func()
}

type ConfigWatcher struct {
	Path    string
	OnEvent func()
}

func (fw *FileWatcher) Watch() {
	if f, err := os.Stat(fw.Path); !f.IsDir() {
		fw.Path = filepath.Base(fw.Path)
	} else if err != nil {
		Log.Critical(err.Error())
		return
	}
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Create, watcher.Write)
	go func() {
		for {
			select {
			case event := <-w.Event:
				Log.Info(event.String())
				fw.OnEvent()
				if fw.Notify {
					fw.notifyComunicato(event)
				}
			case err := <-w.Error:
				Log.Error(err.Error())
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add(fw.Path); err != nil {
		Log.Error(err.Error())
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		Log.Error(err.Error())
	}
}

func (fw FileWatcher) notifyComunicato(event watcher.Event) {
	var (
		tipo    = ""
		dirPath = ""
	)
	dirPath = strings.Replace(event.Path, event.FileInfo.Name(), "", -1)
	if strings.Contains(dirPath, "genitori") {
		tipo = com.TipoGenitori
	} else if strings.Contains(dirPath, "docenti") {
		tipo = com.TipoDocenti
	} else if strings.Contains(dirPath, "studenti") {
		tipo = com.TipoStudenti
	}
	NotifyComunicato(event.FileInfo.Name(), tipo)
}

func (cw *WebContentWatcher) Watch() {
	if f, err := os.Stat(cw.Path); !f.IsDir() {
		cw.Path = filepath.Base(cw.Path)
	} else if err != nil {
		Log.Critical(err.Error())
		return
	}
	w := watcher.New()
	w.SetMaxEvents(1)
	go func() {
		for {
			select {
			case event := <-w.Event:
				Log.Info(event.String())
				cw.OnEvent()
			case err := <-w.Error:
				Log.Error(err.Error())
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add(cw.Path); err != nil {
		Log.Error(err.Error())
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		Log.Error(err.Error())
	}
}

func (cfgw *ConfigWatcher) Watch() {
	w := watcher.New()
	w.SetMaxEvents(1)
	go func() {
		for {
			select {
			case event := <-w.Event:
				if event.FileInfo.Name() == config.GetConfigFilename() {
					Log.Info(event.String())
					cfgw.OnEvent()
				}
			case err := <-w.Error:
				Log.Error(err.Error())
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add(cfgw.Path); err != nil {
		Log.Error(err.Error())
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		Log.Error(err.Error())
	}
}
