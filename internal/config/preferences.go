package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type config struct {
	General general  `json:"generali" toml:"generali"`
	HTTP    connhttp `json:"http" toml:"http"`
	DB      db       `json:"db" toml:"db"`
	Dirs    dirs     `json:"cartelle" toml:"cartelle"`
	Log     logPrefs `json:"logging" toml:"logging"`
}

type general struct {
	Notifications  bool `json:"notifiche`
	RestartOnPanic bool `json:"riavvio_automatico" toml:"riavvio_automatico"`
}

type connhttp struct {
	Port string `json:"porta" toml:"porta"`
}

type dirs struct {
	HTML     string `json:"html" toml:"html"`
	Genitori string `json:"comunicati_genitori" toml:"comunicati_genitori"`
	Studenti string `json:"comunicati_studenti" toml:"comunicati_studenti"`
	Docenti  string `json:"comunicati_docenti" toml:"comunicati_docenti"`
	Progetti string `json:"progetti" toml:"progetti"`
	Orario   string `json:"orario" toml:"orario"`
}

type db struct {
	Schema  string `json:"database" toml:"database"`
	Timeout int64  `json:"timeout" toml:"timeout"`
}

type logPrefs struct {
	Enabled  bool   `json:"abilitato" toml:"abilitato"`
	LogFile  string `json:"file_log" toml:"file_log"`
	LogLevel string `json:"livello_log" toml:"livello_log"`
}

var (
	currentFilePath     = ""
	currentFullFilePath = ""
	currentFiletype     = "none"

	filetypes = []string{
		".json",
		".toml",
	}

	preferences = defaultPrefs

	defaultPrefs = config{
		general{
			Notifications:  false,
			RestartOnPanic: false,
		},
		connhttp{
			Port: ":8080",
		},
		db{
			Schema:  "sitoliceo",
			Timeout: 10,
		},
		dirs{
			HTML:     "./static",
			Genitori: "",
			Studenti: "",
			Docenti:  "",
			Progetti: "",
			Orario:   "./orario.xml",
		},
		logPrefs{
			Enabled:  true,
			LogFile:  "./webapi.log",
			LogLevel: "warning",
		},
	}
)

func LoadPrefs(path string) error {
	preferences = defaultPrefs
	absPath, err := filepath.Abs(path)
	currentFullFilePath = absPath
	currentFilePath = filepath.Dir(absPath)
	if err != nil {
		return err
	}
	for _, filetype := range filetypes {
		if strings.HasSuffix(path, filetype) {
			currentFiletype = filetype
		}
	}
	if currentFiletype == "none" {
		raw, errIO := ioutil.ReadFile(absPath)
		if errIO != nil {
			return errIO
		}
		if !json.Valid(raw) {
			if _, err := toml.DecodeFile(absPath, &config{}); err != nil {
				currentFiletype = "none"
			} else {
				currentFiletype = ".toml"
			}
		}
		if _, err := toml.DecodeFile(absPath, &config{}); err != nil {
			currentFiletype = "none"
		}
	}

	switch currentFiletype {

	case ".json":
		raw, errIO := ioutil.ReadFile(absPath)
		if errIO != nil {
			return errIO
		}
		if errJson := json.Unmarshal(raw, &preferences); errJson != nil {
			return errJson
		}

	case ".toml":
		if _, err := toml.DecodeFile(absPath, &preferences); err != nil {
			return err
		}

	case "none":
		return errors.New("preferences: File configurazione non valido: " + path)
	}

	formatPrefs()
	return nil
}

func formatPrefs() {
	//HTTP
	if !strings.HasPrefix(preferences.HTTP.Port, ":") {
		preferences.HTTP.Port = ":" + preferences.HTTP.Port
	}

	// Dirs
	if _, err := os.Stat(preferences.Dirs.HTML); os.IsNotExist(err) || preferences.Dirs.HTML == "" {
		fmt.Println("Cartella contenuti HTML non specificata")
	}

	// Logging
	if preferences.Log.Enabled {
		if preferences.Log.LogFile == "" {
			preferences.Log.LogFile = "./webapi.log"
		}

		switch preferences.Log.LogLevel {
		case "verbose":
			break

		case "error":
			break

		case "warning":
			break

		default:
			fmt.Println(`formatPrefs: Il livello del log può essere "verbose", "warning" o "error", ignoro opzione, default: "warning"`)
		}
	}
}

func ReloadPrefs() {
	err := LoadPrefs(currentFullFilePath)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *config {
	return &preferences
}

func GetConfigPath() string {
	return currentFilePath
}

func GetConfigFilename() string {
	return filepath.Base(currentFullFilePath)
}
