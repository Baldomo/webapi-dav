package main

import (
	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	formatLong  = logging.MustStringFormatter("[%{level}%{time: 15:04:05.999} %{shortfile}] %{message}")
	formatShort = logging.MustStringFormatter("[%{time:0102 15:04:05.999}] %{message}")
	fileBackend logging.Backend

	Log = logging.MustGetLogger("webapi-dav")
)

func EventLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		Log.Info(
			r.Method + " " +
				r.RequestURI + " " +
				name + " " +
				time.Since(start).String() + " " +
				r.UserAgent(),
		)
	})
}

func InitLogger(before func()) {
	var backendFileFormatted logging.LeveledBackend

	if GetConfig().Log.Enabled {
		l := &lumberjack.Logger{
			Filename: GetConfig().Log.LogFile,
			MaxSize:  10,
			MaxAge:   30,
		}
		fileBackend = logging.NewLogBackend(l, "", 0)

		sign := make(chan os.Signal, 1)
		signal.Notify(sign, os.Interrupt, syscall.SIGTERM)
		go func() {
			for {
				<-sign
				Log.Warning("Salvataggio log...")
				l.Close()
			}
		}()
	} else {
		fileBackend = logging.NewLogBackend(ioutil.Discard, "", 0)
	}

	switch strings.ToLower(GetConfig().Log.LogLevel) {
	case "verbose":
		backendFileFormatted = logging.AddModuleLevel(logging.NewBackendFormatter(fileBackend, formatLong))
		backendFileFormatted.SetLevel(logging.INFO, "")

	case "warning":
		backendFileFormatted = logging.AddModuleLevel(logging.NewBackendFormatter(fileBackend, formatLong))
		backendFileFormatted.SetLevel(logging.WARNING, "")

	case "error":
		backendFileFormatted = logging.AddModuleLevel(logging.NewBackendFormatter(fileBackend, formatLong))
		backendFileFormatted.SetLevel(logging.ERROR, "")

	}

	logging.SetBackend(backendFileFormatted)

	before()
}
