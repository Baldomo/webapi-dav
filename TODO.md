# TODO
- Finire `strings.go` (map codes) + 404.html (static)
- ***WIP* - Eventualmente convertire a [FastCGI](https://github.com/bsingr/golang-apache-fastcgi/blob/master/examples/vanilla/hello_world.go) per Apache
    --> vedere [installazione di mod_fcgi](https://github.com/FastCGI-Archives/mod_fastcgi/blob/master/INSTALL.AP2.md)**
- Implementare progetti
- Aggiungere supporto a token/eTAGs (con **mooooolta** calma)

## Risolti
- ~~Data race in `temp.Execute(w, GetAllOps())`~~
- ~~Aggiungere timeout richieste per evitare flooding~~
- ~~Aggiungere elementi di template a index.html~~
- ~~Mettere a posto `FileWatcher{}.Watch()` (funziona ma non si ricaricano le liste comunicati per qualche ragione)~~
- ~~Aggiungere supporto ad altri tipi di Accept (es. application/xml)~~
- ~~Buttare via Logln e finire InitLogger con [go-logging](https://godoc.org/github.com/op/go-logging)~~
