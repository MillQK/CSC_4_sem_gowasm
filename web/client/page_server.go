// A basic HTTP server.
// By default, it serves the current working directory on port 18080.
package main

import (
	"flag"
	"log"
	"mime"
	"net/http"
)

var (
	listen = flag.String("listen", ":18080", "listen address")
	dir    = flag.String("dir", "./web", "directory to serve")
)

func setup() error {
	if err := mime.AddExtensionType(".wasm", "application/wasm"); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	if err := setup(); err != nil {
		log.Fatalf("Can't setup server, reason: %s", err.Error())
	}
	log.Printf("listening on %q...", *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}
