# webapi-dav
Questa è la documentazione ufficiale di Da Vinci API, contenente utilizzo e configurazione del webserver.
Ultima versione: v0.1.0

## Indice
* [Configurazione](#configurazione)
  * [Generali](#generali)
  * [Connessione](#connessione)
  * [Cartelle](#cartelle)
  * [Logging](#logging)
  * [Esempio](#esempio-configurazione)

## Configurazione
Chiavi e parametri:

#### Generali
| Nome                 | Tipo    | Valori             | Descrizione |
|----------------------|---------|--------------------|-------------|
| `riavvio_automatico` | boolean | `true, false`      | Ancora da implementare, non funzionante. |
| `index_html`         | string  | "/percorso/"       | Indica il percorso del file HTML da usare come pagina principale, indirizzo /api |

#### Connessione
| Nome                 | Tipo     | Valori         | Descrizione |
|----------------------|----------|----------------|-------------|
| `porta`              | string   | "numero porta" | Indica la porta che il server userà per le connessioni in entrata. **Non necessaria** se `apache_cgi` è `true` |
| `apache_cgi`         | boolean  | `true, false`  | Avvia il server in modalità [FastCGI](https://httpd.apache.org/docs/2.4/mod/mod_proxy_fcgi.html) (richiede attivazione di mod_proxy_fcgi in Apache) |

#### Cartelle
| Nome                  | Tipo    | Valori             | Descrizione |
|-----------------------|---------|--------------------|-------------|
| `comunicati_genitori` | string  | "/percorso/"       | Percorso della cartella comunicati genitori |
| `comunicati_studenti` | string  | "/percorso/"       | Percorso della cartella comunicati studenti |
| `comunicati_docenti`  | string  | "/percorso/"       | Percorso della cartella comunicati docenti |
| `progetti`            | string  | "/percorso/"       | Percorso della cartella progetti (non ancora implementato) |

#### Logging
| Nome               | Tipo    | Valori                    | Descrizione |
|--------------------|---------|---------------------------|-------------|
| `log_in_terminale` | boolean | `true, false`             |             |
| `salva_su_file`    | boolean | `true, false`             |             |
| `file_log`         | string  | "/percorso/"              |             |
| `livello_log`      | string  | "verbose, error, warning" |             |

#### Esempio configurazione
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