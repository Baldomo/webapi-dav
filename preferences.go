package main

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type config struct {
	General general   `json:"generali" toml:"generali"`
	HTTPS   connhttps `json:"https" toml:"https"`
	HTTP    connhttp  `json:"http" toml:"http"`
	Dirs    dirs      `json:"cartelle" toml:"cartelle"`
	Log     logPrefs  `json:"logging" toml:"logging"`
}

type general struct {
	RestartOnPanic bool   `json:"riavvio_automatico" toml:"riavvio_automatico"`
	IndexHTML      string `json:"index_html" toml:"index_html"`
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
	Genitori string `json:"comunicati_genitori" toml:"comunicati_genitori"`
	Studenti string `json:"comunicati_studenti" toml:"comunicati_studenti"`
	Docenti  string `json:"comunicati_docenti" toml:"comunicati_docenti"`
	Progetti string `json:"progetti" toml:"progetti"`
}

type logPrefs struct {
	WriteStd  bool   `json:"log_in_terminale" toml:"log_in_terminale"`
	WriteFile bool   `json:"salva_su_file" toml:"salva_su_file"`
	LogFile   string `json:"file_log" toml:"file_log"`
	LogLevel  string `json:"livello_log" toml:"livello_log"`
}

var (
	currentFiletype = "none"

	filetypes = []string{
		".json",
		".toml",
	}

	preferences = defaultPrefs

	defaultPrefs = config{
		general{
			false,
			"./src/static/index.html",
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
		dirs{
			"",
			"",
			"",
			"",
		},
		logPrefs{
			true,
			false,
			"./webapi.log",
			"verbose",
		},
	}
)

func LoadPrefs(path string) error {
	absPath, err := filepath.Abs(path)
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
		return Error("preferences: ", "File configurazione non valido: %s", path)
	}

	formatPrefs()
	return nil
}

func formatPrefs() {
	if !strings.HasPrefix(preferences.HTTPS.Port, ":") {
		preferences.HTTPS.Port = ":" + preferences.HTTPS.Port
	}
	if !strings.HasPrefix(preferences.HTTP.Port, ":") {
		preferences.HTTP.Port = ":" + preferences.HTTP.Port
	}
	if preferences.HTTPS.Enabled {
		if preferences.HTTPS.Cert == "" {
			Log.Fatal("Certificato non specificato!")
		}
		if preferences.HTTPS.Key == "" {
			Log.Fatal("Chiave non specificata!")
		}

	}
	if preferences.Log.WriteStd || preferences.Log.WriteFile {
		switch preferences.Log.LogLevel {
		case "verbose":
			break

		case "error":
			break

		case "warning":
			break

		default:
			Log.Warning(`formatPrefs: Il livello del log può essere "verbose", "warning" o "error", ignoro opzione, default: "warning"`)
			preferences.Log.LogLevel = "warning"
		}
	}
}

func GetConfig() *config {
	return &preferences
}
