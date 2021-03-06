<p align="center">
  <img src="docs/logo.png" width="100" />
</p>
<h1 align="center">webapi-dav</h1>

Questa è la documentazione ufficiale di Da Vinci API (`webapi-dav`), contenente utilizzo e configurazione del webserver.

## Indice
- [Indice](#indice)
- [Compilazione](#compilazione)
    - [Esempio di workflow (debug)](#esempio-di-workflow-debug)
    - [Esempio di workflow (release)](#esempio-di-workflow-release)
- [Esecuzione](#esecuzione)
    - [Variabili d'ambiente](#variabili-dambiente)
    - [Apache](#apache)
- [Configurazione](#configurazione)
    - [Generali](#generali)
    - [Connessione](#connessione)
      - [HTTP](#http)
    - [Cartelle](#cartelle)
    - [Database](#database)
    - [Logging](#logging)
      - [Note:](#note)
- [Esempi configurazione](#esempi-configurazione)
    - [TOML](#toml)
    - [JSON](#json)
- [Comunicati](#comunicati)
- [Altro](#altro)

## Compilazione
Tutte le funzionalità utili a compilare/testare il progetto sono incluse nel file `build.go`. Per ottenere una vista generale delle opzioni disponibili, si veda `go run build.go --help` o il contenuto del file. Di seguito sono riportati alcuni esempi di comandi per compilare/testare `webapi-dav`.

**Nota:** la sintassi dei comandi è sempre `go run build.go <opzioni> <comando>`

#### Esempio di workflow (debug)
Compilazione del progetto (senza creazioni di archivi compressi, si suppone ambiente Linux):
```
go run build.go -os linux -fast build
```

Creazione dell'ambiente di test runtime ed esecuzione del server (vengono copiati i file di configurazione e creati una serie di finti file PDF per simulare i comunicati):
```
go run build.go -run deploy
```

Rimozione dell'ambiente creato:
```
go run build.go clean
```

#### Esempio di workflow (release)
Esecuzione dei test integrati:
```
go run build.go test
```

Compilazione del progetto (si suppone ambiente Linux):
```
go run build.go -os windows build
```

## Esecuzione
L'eseguibile accetta un solo parametro all'esecuzione: `-config file.{json/toml}`,
per specificare manualmente il percorso del file di configurazione. Anche se il file non avesse
l'estensione corretta, il programma cercherà di interpretarlo e, se valido,
lo utilizzerà.

#### Variabili d'ambiente
`webapi-dav` utilizza varie variabili d'ambiente per accedere a credenziali sensibili senza doverle memorizzare in un file.

- `WEBAPI_DB_USER`: username per l'accesso al database degli eventi MySQL
- `WEBAPI_DB_PWD`: password per l'accesso al database degli eventi MySQL
- `WEBAPI_FCM_KEY`: chiave API per il servizio di Firebase Cloud Messaging (notifiche comunicati)

In caso le variabili `WEBAPI_DB_*` non siano inizializzate, la connessione al database non viene eseguita e qualunque query viene semplicemente terminata, per cui ogni richiesta REST riceve come risultato un array JSON vuoto. Per quanto riguarda `WEBAPI_FCM_KEY`, se nulla le notifiche non vengono inviate.

#### Apache
Per utilizzare `webapi-dav` con Apache è necessario usare `mod_proxy`; moduli richiesti:
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


## Configurazione
Sono accettate configurazioni in formato `toml` e `json`. Gli [esempi](#esempi-configurazione)
contengono la configurazione di default in entrambi i formati.
Chiavi e parametri:

#### Generali
| Nome                 | Tipo    | Valori             | Descrizione |
|----------------------|---------|--------------------|-------------|
| `fqdn_sito`          | string  | URL                | Dominio base del server (es. `liceodavinci.tv`) |
| `notifiche`          | boolean | `true, false`      | Attiva il servizio di notifiche tramite Firebase |
| `riavvio_automatico` | boolean | `true, false`      | Ancora da implementare, non funzionante |

---

#### Connessione

##### HTTP
| Nome                 | Tipo     | Valori         | Descrizione |
|----------------------|----------|----------------|-------------|
| `porta`              | string   | "numero porta" | Indica la porta che il server userà per le connessioni HTTP in entrata |

---

#### Cartelle
| Nome                  | Tipo    | Valori             | Descrizione |
|-----------------------|---------|--------------------|-------------|
| `html`                | string  | "/percorso/"       | Percorso della cartella che contiene i file HTML per le pagine web |
| `path_comunicati`     | string  | "/percorso/"       | Percorso della cartella comunicati |
| `comunicati_genitori` | string  | "/percorso/"       | Percorso della sottocartella comunicati genitori |
| `comunicati_studenti` | string  | "/percorso/"       | Percorso della sottocartella comunicati studenti |
| `comunicati_docenti`  | string  | "/percorso/"       | Percorso della sottocartella comunicati docenti |
| `progetti`            | string  | "/percorso/"       | Percorso della cartella progetti (**non ancora implementato**) |
| `orario`              | string  | "/percorso/"       | Percorso del file esportato dal gestionale dell'orario, contenente la tabella Attività in XML |

---

#### Database
| Nome       | Tipo    | Valori             | Descrizione |
|------------|---------|--------------------|-------------|
| `database` | string  | `true, false`      | Nome del database Joomla contenente l'agenda |
| `timeout`  | integer | numero intero      | Timeout per connessione al database |

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
[generali]
fqdn_sito = "liceodavinci.tv"
notifiche = false

[autenticazione]
chiave_firma = "secret"

[http]
porta = ":8080"

[db]
database = "sitoliceo"

[cartelle]
html = "docs"
path_comunicati = "/sitoLiceo/images/comunicati/"
comunicati_genitori = "comunicati-genitori"
comunicati_studenti = "comunicati-studenti"
comunicati_docenti = "comunicati-docenti"
orario = "orario.xml"

[logging]
abilitato = true
file_log = "./run.log"
livello_log = "verbose"
```

#### JSON
Esempio di `config.json` (in quella di default le cartelle non sono specificate):
```json
{
  "generali": {
    "fqdn_sito": "liceodavinci.tv",
    "notifiche": false
  },
  "autenticazione": {
    "chiave_firma": "secret"
  },
  "http": {
    "porta": ":8080"
  },
  "db": {
    "database": "sitoliceo"
  },
  "cartelle": {
    "html": "docs",
    "path_comunicati": "/sitoLiceo/images/comunicati/",
    "comunicati_genitori": "comunicati-genitori",
    "comunicati_studenti": "comunicati-studenti",
    "comunicati_docenti": "comunicati-docenti",
    "orario": "orario.xml"
  },
  "logging": {
    "abilitato": true,
    "file_log": "./run.log",
    "livello_log": "verbose"
  }
}
```

## Comunicati
Qualunque cartella contenente file può essere esposta nella configurazione: la API
terrà in memoria una lista dei comunicati secondo la struttura:
```go
type Comunicato struct {
    Nome  string    // nome del file
    Data  time.Time // data di ultima modifica
    Tipo  string    // "genitori", "studenti" o "docenti"
    URL   string    // link diretto al PDF
}
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