# MILESTONE
- ![0.2.0](http://progressed.io/bar/25?title=v0.2.0)
- [![0.1.0](http://progressed.io/bar/100?title=v0.1.0)](https://bitbucket.org/Baldomo/webapi-dav/commits/03d5f82f2d93)

# TODO
- Aggiungere versione e altri elementi variabili a `strings.go`
     - Eventualmente aggiungere una template di base con relativa `struct`

## In attesa
- Integrazione [Docker](https://blog.golang.org/docker)
- Implementare progetti
- Aggiungere supporto a token/eTAGs (con **mooooolta** calma)

## Risolti
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
