# webapi-dav
Questa è la documentazione ufficiale di Da Vinci API, contenente utilizzo e configurazione del webserver.
Ultima versione: v0.1.0

## Indice
* [Configurazione](#configurazione)
  * [Generali](#generali)
  * [Connessione](#connessione)
  * [Cartelle](#cartelle)
  * [Logging](#logging)
  * [Esempi](#esempi-configurazione)
    * [TOML](#toml)
    * [JSON](#json)


## Esecuzione
L'eseguibile accetta un solo parametro all'esecuzione: `-config file.{json/toml}`,
per specificare manualmente il percorso del file di configurazione

## Configurazione
Sono accettate configurazioni in formato `toml` e `json`. Gli [esempi](#esempi-configurazione)
contengono la configurazione di default in entrambi i formati.
Chiavi e parametri:

#### Generali
| Nome                 | Tipo    | Valori             | Descrizione |
|----------------------|---------|--------------------|-------------|
| `riavvio_automatico` | boolean | `true, false`      | Ancora da implementare, non funzionante. |
| `index_html`         | string  | "/percorso/"       | Indica il percorso del file HTML da usare come pagina principale, indirizzo /api |

---

#### Connessione
| Nome                 | Tipo     | Valori         | Descrizione |
|----------------------|----------|----------------|-------------|
| `porta`              | string   | "numero porta" | Indica la porta che il server userà per le connessioni in entrata. **Non necessaria** se `apache_cgi` è `true` |
| `apache_cgi`         | boolean  | `true, false`  | Avvia il server in modalità [FastCGI](https://httpd.apache.org/docs/2.4/mod/mod_proxy_fcgi.html) (richiede attivazione di mod_proxy_fcgi in Apache) |

---

#### Cartelle
| Nome                  | Tipo    | Valori             | Descrizione |
|-----------------------|---------|--------------------|-------------|
| `comunicati_genitori` | string  | "/percorso/"       | Percorso della cartella comunicati genitori |
| `comunicati_studenti` | string  | "/percorso/"       | Percorso della cartella comunicati studenti |
| `comunicati_docenti`  | string  | "/percorso/"       | Percorso della cartella comunicati docenti |
| `progetti`            | string  | "/percorso/"       | Percorso della cartella progetti (non ancora implementato) |

---

#### Logging
| Nome               | Tipo    | Valori                    | Descrizione |
|--------------------|---------|---------------------------|-------------|
| `log_in_terminale` | boolean | `true, false`             | Definisce se stampare i log sul terminale |
| `salva_su_file`    | boolean | `true, false`             | Definisce se stampare i log su un file |
| `file_log`         | string  | "/percorso/"              | Percorso del file di log, se non esiste sarà creato. |
| `livello_log`      | string  | "verbose, error, warning" | Definisce il livello di verbosità dei log, con "verbose" il massimo (ogni evento sarà tracciato) e "error" il minimo (solo gli errori interni saranno tracciati) |

##### Note:
Se sia `log_in_terminale` che `salva_su_file` sono `false`, nessun evento sarà tracciato

---

## Esempi configurazione
#### TOML
Esempio di `config.toml` (configurazione di default):

```toml
[generali]
riavvio_automatico = true
index_html = "./static/index.html"

[connessione]
porta = "8080"
apache_cgi = false

[cartelle]
comunicati_genitori = "./comunicati-genitori"
comunicati_studenti = "./comunicati-studenti"
comunicati_docenti = "./comunicati-docenti"
progetti = "./progetti"

[logging]
log_in_terminale = true
salva_su_file = false
file_log = "./webapi.log"
livello_log = "verbose"
```

#### JSON
Esempio di `config.json` (configurazione di default):
```json
{
  "generali": {
    "riavvio_automatico": true,
    "index_html": "./static/index.html"
  },
  "connessione": {
    "porta": "8080",
    "apache_cgi": false
  },
  "cartelle": {
    "comunicati_genitori": "./comunicati-genitori",
    "comunicati_studenti": "./comunicati-studenti",
    "comunicati_docenti": "./comunicati-docenti",
    "progetti": "./progetti"
  },
  "logging": {
    "log_in_terminale": true,
    "salva_su_file": false,
    "file_log": "./webapi.log",
    "livello_log": "verbose"
  }
}
```