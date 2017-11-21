package main

import (
	"github.com/op/go-logging"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	formatLong       = logging.MustStringFormatter("%{color}[%{level}%{time: 15:04:05.999999} %{shortfile}] %{message} %{color:reset}")
	formatShort      = logging.MustStringFormatter("[%{time:0102 15:04:05.999999}] %{message}")
	stdBackend       logging.Backend
	fileBackend      logging.Backend
	compiledBackends []logging.Backend

	Log = logging.MustGetLogger("webapi-dav")
)

func EventLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		Log.Info(
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func InitLogger() {
	if runtime.GOOS == "windows" {
		logging.MustStringFormatter("[%{level:-1s}%{time: 15:04:05.999999} %{shortfile}] %{message}")
	}
	if GetConfig().Log.WriteStd {
		stdBackend = logging.NewLogBackend(os.Stdout, "", 0)

	} else {
		stdBackend = logging.NewLogBackend(ioutil.Discard, "", 0)
	}
	if GetConfig().Log.WriteFile {
		file, err := os.Create(GetConfig().Log.LogFile)
		if err != nil {
			panic(err)
		}
		fileBackend = logging.NewLogBackend(file, "", 0)
	} else {
		fileBackend = logging.NewLogBackend(ioutil.Discard, "", 0)
	}
	if !GetConfig().Log.WriteStd && !GetConfig().Log.WriteFile {
		defer Log.Info("-* Logger avviato in modalità silenziosa *-")
	}

	switch strings.ToLower(GetConfig().Log.LogLevel) {
	case "verbose":
		backendStdFormatted := logging.AddModuleLevel(logging.NewBackendFormatter(stdBackend, formatLong))
		backendStdFormatted.SetLevel(logging.INFO, "")
		backendFileFormatted := logging.AddModuleLevel(logging.NewBackendFormatter(fileBackend, formatLong))
		backendFileFormatted.SetLevel(logging.INFO, "")
		compiledBackends = append(compiledBackends, backendStdFormatted)

	case "warning":
		backendStdFormatted := logging.AddModuleLevel(logging.NewBackendFormatter(stdBackend, formatLong))
		backendStdFormatted.SetLevel(logging.WARNING, "")
		backendFileFormatted := logging.AddModuleLevel(logging.NewBackendFormatter(fileBackend, formatLong))
		backendFileFormatted.SetLevel(logging.WARNING, "")
		compiledBackends = append(compiledBackends, backendStdFormatted)

	case "error":
		backendStdFormatted := logging.AddModuleLevel(logging.NewBackendFormatter(stdBackend, formatLong))
		backendStdFormatted.SetLevel(logging.ERROR, "")
		backendFileFormatted := logging.AddModuleLevel(logging.NewBackendFormatter(fileBackend, formatLong))
		backendFileFormatted.SetLevel(logging.ERROR, "")
		compiledBackends = append(compiledBackends, backendStdFormatted)

	}

	logging.SetBackend(compiledBackends...)
}
