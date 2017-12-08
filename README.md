# webapi-dav
Questa è la documentazione ufficiale di Da Vinci API, contenente utilizzo e configurazione del webserver.
Ultima versione: v0.2.0
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
L'unico modo (al momento) di poter eseguire il webserver è tramite mod_proxy: in `httpd.conf` usare moduli:
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

VirtualHost per redirezionamento richieste, aggiungere sempre a `httpd.conf`:
```apache
<VirtualHost *:80>
    ServerName localhost/api

    ProxyRequests Off
    ProxyVia Off

    ProxyPass /api http://127.0.0.1:8080/api
    ProxyPassReverse /api http://127.0.0.1:8080/api
</VirtualHost>
```
La modalità [FastCGI](https://httpd.apache.org/docs/2.4/mod/mod_proxy_fcgi.html) è in sviluppo, insieme all'integrazione del sistema di gestione servizi di Windows.

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
| `apache_cgi`         | boolean  | `true, false`  | Avvia il server in modalità [FastCGI](https://httpd.apache.org/docs/2.4/mod/mod_proxy_fcgi.html) (richiede attivazione di mod_fcgi in Apache, in sviluppo) |
| `https`              | boolean  | `true, false`  | Avvia il server in modalità HTTPS - richiede percorsi di certificato e chiave privata del certificato |
| `certificato`        | string   | "/percorso/"   | Indica il percorso del file `.crt` del certificato firmato |
| `chiave`             | string   | "/percorso/"   | Indica il percorso del file `.key` contenente la chiave con cui è stato firmato il certificato |

##### Note:
Impostare HTTPS a `true` sovrascriverà la porta a `:443`, standard per HTTPS

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
Esempio di `config.toml` (in quella di default le cartelle non sono specificate):

```toml
[generali]
riavvio_automatico = true
index_html = "static/index.html"

[connessione]
porta = "8080"
apache_cgi = false
https = false
certificato = "server.crt"
chiave = "server.key"

[cartelle]
comunicati_genitori = "comunicati-genitori"
comunicati_studenti = "comunicati-studenti"
comunicati_docenti = "comunicati-docenti"
progetti = "progetti"

[logging]
log_in_terminale = true
salva_su_file = false
file_log = "webapi.log"
livello_log = "verbose"
```

#### JSON
Esempio di `config.json` (in quella di default le cartelle non sono specificate):
```json
{
  "generali": {
    "riavvio_automatico": true,
    "index_html": "static/index.html"
  },

  "connessione": {
    "porta": "8080",
    "apache_cgi": false,
    "https": false,
    "certificato": "server.crt",
    "chiave": "server.key"
  },

  "cartelle": {
    "comunicati_genitori": "comunicati-genitori",
    "comunicati_studenti": "comunicati-studenti",
    "comunicati_docenti": "comunicati-docenti",
    "progetti": "progetti"
  },

  "logging": {
    "log_in_terminale": true,
    "salva_su_file": false,
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
Esempio di risposta di singolo comunicato in XML:
```xml
<Comunicato>
    <Nome>177_corsa campestre istituto.pdf</Nome>
    <Data>2017-11-26T10:30:49.272711528+01:00</Data>
    <Tipo>studenti</Tipo>
    <URL>http://liceodavinci.tv/sitoLiceo/comunicati/comunicati-studenti/...</URL>
</Comunicato>
```

## Progetti
Ancora da implementare