# MILESTONE
- ![0.1.0](http://progressed.io/bar/100?title=v0.1.0)

# TODO
- Refactor generale del progetto
    - Rimuovere cartelle di test comunicati e fare uno script per generare file nel Makefile
- Aggiungere versione e altri elementi variabili a `strings.go`

## In attesa
- Integrazione [Docker](https://blog.golang.org/docker)
- Implementare progetti
- Aggiungere supporto a token/eTAGs (con **mooooolta** calma)

## Risolti
- ~~Finire 404.html (static)~~
- ~~Data race in `temp.Execute(w, GetAllOps())`~~
- ~~Eventualmente aggiungere build per [FastCGI](https://github.com/bsingr/golang-apache-fastcgi/blob/master/examples/vanilla/hello_world.go) per Apache
    --> vedere [installazione di mod_fcgi](https://github.com/FastCGI-Archives/mod_fastcgi/blob/master/INSTALL.AP2.md)~~
- ~~Aggiungere timeout richieste per evitare flooding~~
- ~~Aggiungere elementi di template a index.html~~
- ~~Mettere a posto `FileWatcher{}.Watch()` (funziona ma non si ricaricano le liste comunicati per qualche ragione)~~
- ~~Aggiungere supporto ad altri tipi di Accept (es. application/xml)~~
- ~~Buttare via Logln e finire InitLogger con [go-logging](https://godoc.org/github.com/op/go-logging)~~
