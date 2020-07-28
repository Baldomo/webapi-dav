<p align="center">
  <img src="docs/logo.png" width="100" />
</p>
<h1 align="center">webapi-dav</h1>

Questa è la documentazione ufficiale di Da Vinci API (`webapi-dav`), contenente utilizzo e configurazione del webserver.

## Indice
* [Compilazione](#compilazione)
* [Esecuzione](#esecuzione)
* [Configurazione](#configurazione)
    * [Apache](#apache)
    * [Generali](#generali)
    * [Connessione](#connessione)
    * [Cartelle](#cartelle)
    * [Logging](#logging)
    * [Esempi](#esempi-configurazione)
        * [TOML](#toml)
        * [JSON](#json)
* [Comunicati](#comunicati)
* ~~[Progetti](#progetti)~~

## Compilazione
Tutte le funzionalità utili a compilare/testare il progetto sono incluse nel file `build.go`. Per ottenere una vista generale delle opzioni disponibili, si veda `go run build.go --help` o il contenuto del file. Di seguito sono riportati alcuni esempi di comandi per compilare/testare `webapi-dav`.

#### Esempio di workflow (debug)
Compilazione del progetto (senza creazioni di archivi compressi, si suppone ambiente Linux):
```
go run build.go build -os linux -fast
```

Creazione dell'ambiente di test runtime ed esecuzione del server (vengono copiati i file di configurazione e creati una serie di finti file PDF per simulare i comunicati):
```
go run build.go deploy -run
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
go run build.go build -os windows
```

## Esecuzione
L'eseguibile accetta un solo parametro all'esecuzione: `-config file.{json/toml}`,
per specificare manualmente il percorso del file di configurazione. Anche se il file non avesse
l'estensione corretta, il programma cercherà di interpretarlo e, se valido,
lo utilizzerà.

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


## Configurazione
Sono accettate configurazioni in formato `toml` e `json`. Gli [esempi](#esempi-configurazione)
contengono la configurazione di default in entrambi i formati.
Chiavi e parametri:

#### Generali
| Nome                 | Tipo    | Valori             | Descrizione |
|----------------------|---------|--------------------|-------------|
| `riavvio_automatico` | boolean | `true, false`      | Ancora da implementare, non funzionante. |

---

#### Connessione

##### HTTP
| Nome                 | Tipo     | Valori         | Descrizione |
|----------------------|----------|----------------|-------------|
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
[http]
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
  "http": {
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