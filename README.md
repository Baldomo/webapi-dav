# webapi-dav
Questa è la documentazione ufficiale di Da Vinci API, contenente utilizzo e configurazione del webserver.
Ultima versione: v0.5.0
[Copia del README su Github](https://gist.github.com/Baldomo/5dc1db7a46e00f94ef714b7063f7fa3d)

## Indice
* [Esecuzione](#markdown-header-esecuzione)
* [Configurazione](#markdown-header-configurazione)
    * [Apache](#markdown-header-apache)
    * [Generali](#markdown-header-generali)
    * [Connessione](#markdown-header-connessione)
    * [Cartelle](#markdown-header-cartelle)
    * [Logging](#markdown-header-logging)
    * [Esempi](#markdown-header-esempi-configurazione)
        * [TOML](#markdown-header-toml)
        * [JSON](#markdown-header-json)
* [Comunicati](#markdown-header-comunicati)
* ~~[Progetti](#markdown-header-progetti)~~


## Esecuzione
L'eseguibile accetta un solo parametro all'esecuzione: `-config file.{json/toml}`,
per specificare manualmente il percorso del file di configurazione. Anche se il file non avesse
l'estensione corretta, il programma cercherà di interpretarlo e, se valido,
lo utilizzerà.

## Configurazione
Sono accettate configurazioni in formato `toml` e `json`. Gli [esempi](#esempi-configurazione)
contengono la configurazione di default in entrambi i formati.
Chiavi e parametri:

#### Apache
L'unico modo (al momento) di poter eseguire il webserver è tramite mod_proxy; moduli richiesti:
```apache
LoadModule headers_module modules/mod_headers.so
LoadModule proxy_module modules/mod_proxy.so
LoadModule proxy_connect_module modules/mod_proxy_connect.so
LoadModule proxy_html_module modules/mod_proxy_html.so
LoadModule proxy_http_module modules/mod_proxy_http.so
LoadModule slotmem_shm_module modules/mod_slotmem_shm.so
LoadModule socache_shmcb_module modules/mod_socache_shmcb.so
LoadModule xml2enc_module modules/mod_xml2enc.so
```

VirtualHost per redirezionamento richieste HTTP:
```apache
<VirtualHost *:80>
    ServerName localhost/api

    ProxyRequests Off
    ProxyVia Off

    ProxyPass /api http://127.0.0.1:80/api
    ProxyPassReverse /api http://127.0.0.1:80/api
</VirtualHost>
```

VirtualHost per redirezionamento richieste HTTPS:
```apache
<VirtualHost *:443>
    ServerName localhost/api

    ProxyRequests Off
    ProxyVia Off

    ProxyPass /api http://127.0.0.1:443/api
    ProxyPassReverse /api http://127.0.0.1:443/api
</VirtualHost>
```

#### Generali
| Nome                 | Tipo    | Valori             | Descrizione |
|----------------------|---------|--------------------|-------------|
| `riavvio_automatico` | boolean | `true, false`      | Ancora da implementare, non funzionante. |

---

#### Connessione

##### HTTPS
| Nome                 | Tipo     | Valori         | Descrizione |
|----------------------|----------|----------------|-------------|
| `abilitato`          | boolean  | `true, false`  | Avvia il server in modalità HTTPS - richiede percorsi di certificato e chiave privata del certificato |
| `porta`              | string   | "numero porta" | Indica la porta che il server userà per le connessioni HTTPS in entrata. |
| `certificato`        | string   | "/percorso/"   | Indica il percorso del file `.crt` del certificato firmato |
| `chiave`             | string   | "/percorso/"   | Indica il percorso del file `.key` contenente la chiave con cui è stato firmato il certificato |

##### HTTP
| Nome                 | Tipo     | Valori         | Descrizione |
|----------------------|----------|----------------|-------------|
| `abilitato`          | boolean  | `true, false`  | Avvia il server in modalità HTTP |
| `porta`              | string   | "numero porta" | Indica la porta che il server userà per le connessioni HTTP in entrata. |

---

#### Cartelle
| Nome                  | Tipo    | Valori             | Descrizione |
|-----------------------|---------|--------------------|-------------|
| `html`                | string  | "/percorso/"       | Percorso della cartella che contiene i file HTML per le pagine web |
| `comunicati_genitori` | string  | "/percorso/"       | Percorso della cartella comunicati genitori |
| `comunicati_studenti` | string  | "/percorso/"       | Percorso della cartella comunicati studenti |
| `comunicati_docenti`  | string  | "/percorso/"       | Percorso della cartella comunicati docenti |
| `progetti`            | string  | "/percorso/"       | Percorso della cartella progetti (**non ancora implementato**) |
| `orario`              | string  | "/percorso/"       | Percorso del file esportato dal gestionale dell'orario, contenente la tabella Attività in XML |

---

#### Logging
| Nome          | Tipo    | Valori                    | Descrizione |
|---------------|---------|---------------------------|-------------|
| `abilitato`   | boolean | `true, false`             | Avvia il servizio di logging |
| `file_log`    | string  | "/percorso/"              | Percorso del file di log, se non esiste sarà creato. |
| `livello_log` | string  | "verbose, error, warning" | Definisce il livello di verbosità dei log, con "verbose" il massimo (ogni evento sarà tracciato) e "error" il minimo (solo gli errori interni saranno tracciati) |

##### Note:
- Se il logging è disabilitato, nessun evento sarà tracciato
- I log ruotano automaticamente ogni 30 giorni o ogni 5 MB occupati per file, dopodichè saranno
    rinominati con la data corrente e compattati in un archivio zip

---

## Esempi configurazione
#### TOML
Esempio di `config.toml` (in quella di default le cartelle non sono specificate):

```toml
[https]
abilitato = false
porta = ":443"
certificato = "server.crt"
chiave = "server.key"

[http]
abilitato = true
porta = ":8080"

[db]
database = "sitoliceo"

[cartelle]
html = "static/"
comunicati_genitori = "comunicati-genitori"
comunicati_studenti = "comunicati-studenti"
comunicati_docenti = "comunicati-docenti"
orario = "orario.xml"

[logging]
abilitato = true
file_log = "webapi.log"
livello_log = "verbose"
```

#### JSON
Esempio di `config.json` (in quella di default le cartelle non sono specificate):
```json
{
  "https": {
    "abilitato": false,
    "porta": ":443",
    "certificato": "server.crt",
    "chiave": "server.key"
  },
  "http": {
    "abilitato": true,
    "porta": ":8080"
  },
  "db": {
    "database": "sitoliceo",
  },
  "cartelle": {
    "html": "static/",
    "comunicati_genitori": "comunicati-genitori",
    "comunicati_studenti": "comunicati-studenti",
    "comunicati_docenti": "comunicati-docenti",
    "orario": "orario.xml"
  },
  "logging": {
    "abilitato": true,
    "file_log": "webapi.log",
    "livello_log": "verbose"
  }
}
```

## Comunicati
Qualunque cartella contenente file può essere esposta nella configurazione: la API
terrà in memoria una lista dei comunicati secondo la struttura:
```go
Comunicato
    Nome  string    (nome del file)
    Data  time.Time (data di ultima modifica)
    Tipo  string    ("genitori", "studenti" o "docenti")
    URL   string    (link diretto al PDF)
```
Esempio di risposta di singolo comunicato in JSON:
```json
{
    "nome":"177_corsa campestre istituto.pdf",
    "data":"2017-11-26T10:30:49.272711528+01:00",
    "tipo":"studenti",
    "url":"http://liceodavinci.tv/sitoLiceo/comunicati/comunicati-studenti/..."
}
```

## Altro
Version, About, ecc: risposta JSON:
```json
{
  "codice": 200,
  "info": "Leonardo Baldin, v0.3.0, (c) 2017"
}
```