# MILESTONE
- [![0.5.0](http://progressed.io/bar/100?title=v0.5.0)](https://bitbucket.org/Baldomo/webapi-dav/commits/tag/v0.5.0)
- [![0.4.0](http://progressed.io/bar/100?title=v0.4.0)](https://bitbucket.org/Baldomo/webapi-dav/commits/tag/v0.4.0)
- [![0.3.0](http://progressed.io/bar/100?title=v0.3.0)](https://bitbucket.org/Baldomo/webapi-dav/commits/tag/v0.3.0)
- [![0.2.0](http://progressed.io/bar/100?title=v0.2.0)](https://bitbucket.org/Baldomo/webapi-dav/commits/tag/v0.2.0)
- [![0.1.0](http://progressed.io/bar/100?title=v0.1.0)](https://bitbucket.org/Baldomo/webapi-dav/commits/tag/v0.1.0)

# TODO
- Aggiungere test per ogni pkg
- Fixare ReloadPrefs in `preferences.go`

## In attesa
- Implementare progetti

## Risolti
- ~~Delegare credenziali DB a variabile di sistema~~
- ~~Finire interfaccia database~~
- ~~Endpoint orario in `routes.go#95`~~
- ~~Finire implementazione close `startup_windows.go#24` (protocollo custom)~~
    - ~~Evitare forking o altri processi~~
- ~~Rotazione log con Lumberjack~~
- ~~Usare una struct generica per risposte a endpoint /about, /version ecc.~~
- ~~Aggiungere supporto a [HTTPS](https://github.com/denji/golang-tls)~~
- ~~Aggiungere elementi a `strings.go`~~
     - ~~Eventualmente aggiungere una template di base con relativa `struct`~~
- ~~Rimuovere cartelle di test comunicati e fare uno script per generare file nel Makefile~~
- ~~Finire 404.html (static)~~
- ~~Data race in `temp.Execute(w, GetAllOps())`~~
- ~~Eventualmente aggiungere build per [FastCGI](https://github.com/bsingr/golang-apache-fastcgi/blob/master/examples/vanilla/hello_world.go) per Apache
    --> vedere [installazione di mod_fcgi](https://github.com/FastCGI-Archives/mod_fastcgi/blob/master/INSTALL.AP2.md)~~
- ~~Aggiungere timeout richieste per evitare flooding~~
- ~~Aggiungere elementi di template a index.html~~
- ~~Mettere a posto `FileWatcher{}.Watch()` (funziona ma non si ricaricano le liste comunicati per qualche ragione)~~
- ~~Aggiungere supporto ad altri tipi di Accept (es. application/xml)~~
- ~~Buttare via Logln e finire InitLogger con [go-logging](https://godoc.org/github.com/op/go-logging)~~
