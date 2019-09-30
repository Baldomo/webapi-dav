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

	"github.com/Baldomo/webapi-dav/internal/config"
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

func InitLogger(before func()) {
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

	before()
}

func CloseLogger() {
	logger.Warning("Salvataggio log...")
	lumber.Close()
}

func Critical(format string, args ...interface{}) { logger.Critical(format, args) }

func Debug(format string, args ...interface{})  { logger.Debug(format, args) }
func Debugf(format string, args ...interface{}) { logger.Debugf(format, args) }

func Error(format string, args ...interface{})  { logger.Error(format, args) }
func Errorf(format string, args ...interface{}) { logger.Errorf(format, args) }

func Fatal(args ...interface{})                 { logger.Fatal(args...) }
func Fatalf(format string, args ...interface{}) { logger.Fatalf(format, args) }

func Info(format string, args ...interface{})  { logger.Info(format, args) }
func Infof(format string, args ...interface{}) { logger.Infof(format, args) }

func Notice(format string, args ...interface{})  { logger.Notice(format, args) }
func Noticef(format string, args ...interface{}) { logger.Noticef(format, args) }

func Panic(args ...interface{})                 { logger.Panic(args...) }
func Panicf(format string, args ...interface{}) { logger.Panicf(format, args...) }

func Warning(format string, args ...interface{})  { logger.Warning(format, args) }
func Warningf(format string, args ...interface{}) { logger.Warningf(format, args) }
