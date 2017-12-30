package main

import (
	"fmt"
	"github.com/radovskyb/watcher"
	"net/http"
	"strings"
	"time"
)

type FileWatcher struct {
	Path    string
	Store   interface{}
	OnEvent func()
}

type WebContentWatcher struct {
	Path    string
	OnEvent func()
}

type ConfigWatcher struct{}

func RequestMime(header http.Header) string {
	/*if strings.Split(header.Get("Accept"), ",")[0] == "text/html" {
		return "application/json"
	}*/
	return strings.Split(header.Get("Accept"), ",")[0]
}

func Error(origin string, format string, args ...interface{}) error {
	return fmt.Errorf(origin+format, args...)
}

func (fw FileWatcher) Watch() {
	w := watcher.New()
	w.SetMaxEvents(1)
	//w.FilterOps(watcher.Create, watcher.Rename, watcher.Remove, watcher.Write)
	go func() {
		for {
			select {
			case event := <-w.Event:
				Log.Info(event.String())
				fw.OnEvent()
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

func (cw WebContentWatcher) Watch() {
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
