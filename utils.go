package main

import (
	"fmt"
	"github.com/radovskyb/watcher"
	"log"
	"net/http"
	"strings"
	"time"
)

type FileWatcher struct {
	Path    string
	Store   interface{}
	OnEvent func()
}

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
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add(fw.Path); err != nil {
		log.Fatalln(err)
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

//TODO
func CheckAndRecover(err error) {
	if err != nil {
		panic(err)
	}
}
