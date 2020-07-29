package log

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/Baldomo/webapi-dav/pkg/config"
	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	formatLong = logging.MustStringFormatter("[%{level}%{time: 15:04:05.999} %{shortfile}] %{message}")
	//formatShort = logging.MustStringFormatter("[%{time:0102 15:04:05.999}] %{message}")
	fileBackend logging.Backend

	logger = logging.MustGetLogger("webapi-dav")

	lumber = &lumberjack.Logger{
		Filename: config.GetConfig().Log.LogFile,
		MaxSize:  5,
		MaxAge:   30,
		Compress: true,
	}
)

// Middleware http per loggare su disco/terminale le richieste.
// Si noti che non viene rilevato nessun dato sensibile ad eccezione
// dell'indirizzo IP della richiesta e lo User Agent, che sono dati variabili
func EventLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		logger.Info(
			r.Method + " " +
				r.RequestURI + " " +
				name + " " +
				time.Since(start).String() + " " +
				r.RemoteAddr + " " +
				r.UserAgent(),
		)
	})
}

// Inizializza il logger interno: esegue setup di log su disco e/o su terminale
// secondo la configurazione
func InitLogger() {
	var backendFileFormatted logging.LeveledBackend

	if config.GetConfig().Log.Enabled {
		fileBackend = logging.NewLogBackend(lumber, "", 0)

		if runtime.GOOS != "windows" {
			sign := make(chan os.Signal, 1)
			signal.Notify(sign, os.Interrupt, syscall.SIGTERM)
			go func() {
				for {
					<-sign
					logger.Warning("Salvataggio log...")
					lumber.Close()
				}
			}()
		}
	} else {
		fileBackend = logging.NewLogBackend(ioutil.Discard, "", 0)
	}

	switch strings.ToLower(config.GetConfig().Log.LogLevel) {
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
}

// Chiude il logger, termina eventuale scrittura su disco
func CloseLogger() {
	logger.Warning("Salvataggio log...")
	lumber.Close()
}

// Critical logs a message using CRITICAL as log level.
func Critical(format string, args ...interface{}) { logger.Critical(format, args) }

// Debug logs a message using DEBUG as log level.
func Debug(format string, args ...interface{}) { logger.Debug(format, args) }

// Debugf logs a message using DEBUG as log level and a format string.
func Debugf(format string, args ...interface{}) { logger.Debugf(format, args) }

// Error logs a message using ERROR as log level.
func Error(format string, args ...interface{}) { logger.Error(format, args) }

// Errorf logs a message using ERROR as log level and a format string.
func Errorf(format string, args ...interface{}) { logger.Errorf(format, args) }

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
func Fatal(args ...interface{}) { logger.Fatal(args...) }

// Fatalf is equivalent to l.Critical followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) { logger.Fatalf(format, args) }

// Info logs a message using INFO as log level.
func Info(format string, args ...interface{}) { logger.Info(format, args) }

// Infof logs a message using INFO as log level and a format string.
func Infof(format string, args ...interface{}) { logger.Infof(format, args) }

// Notice logs a message using NOTICE as log level.
func Notice(format string, args ...interface{}) { logger.Notice(format, args) }

// Noticef logs a message using NOTICE as log level and a format string.
func Noticef(format string, args ...interface{}) { logger.Noticef(format, args) }

// Panic is equivalent to l.Critical(fmt.Sprint()) followed by a call to panic().
func Panic(args ...interface{}) { logger.Panic(args...) }

// Panicf is equivalent to l.Critical followed by a call to panic().
func Panicf(format string, args ...interface{}) { logger.Panicf(format, args...) }

// Warning logs a message using WARNING as log level.
func Warning(format string, args ...interface{}) { logger.Warning(format, args) }

// Warningf logs a message using WARNING as log level and a format string.
func Warningf(format string, args ...interface{}) { logger.Warningf(format, args) }
