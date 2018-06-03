package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	General general   `json:"generali" toml:"generali"`
	HTTPS   connhttps `json:"https" toml:"https"`
	HTTP    connhttp  `json:"http" toml:"http"`
	DB      db        `json:"db" toml:"db"`
	Dirs    dirs      `json:"cartelle" toml:"cartelle"`
	Log     logPrefs  `json:"logging" toml:"logging"`
}

type general struct {
	RestartOnPanic bool `json:"riavvio_automatico" toml:"riavvio_automatico"`
}

type connhttps struct {
	Enabled bool   `json:"abilitato" toml:"abilitato"`
	Port    string `json:"porta" toml:"porta"`
	Cert    string `json:"certificato" toml:"certificato"`
	Key     string `json:"chiave" toml:"chiave"`
}

type connhttp struct {
	Enabled bool   `json:"abilitato" toml:"abilitato"`
	Port    string `json:"porta" toml:"porta"`
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
	Schema string `json:"database" toml:"database"`
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
			false,
		},
		connhttps{
			false,
			"",
			"",
			"",
		},
		connhttp{
			true,
			":8080",
		},
		db{
			"sitoliceo",
		},
		dirs{
			"./static",
			"",
			"",
			"",
			"",
			"./orario.xml",
		},
		logPrefs{
			true,
			"./webapi.log",
			"warning",
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

	// HTTPS
	if !strings.HasPrefix(preferences.HTTPS.Port, ":") {
		preferences.HTTPS.Port = ":" + preferences.HTTPS.Port
	}
	if preferences.HTTPS.Enabled {
		if preferences.HTTPS.Cert == "" {
			fmt.Println("Certificato non specificato!")
		}
		if preferences.HTTPS.Key == "" {
			fmt.Println("Chiave non specificata!")
		}

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
	LoadPrefs(currentFullFilePath)
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
