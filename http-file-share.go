package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var path = flag.String("path", "", "root for sharing files")
	var listen = flag.String("listen", ":8080", "listen directive")
	flag.Parse()

	if len(*path) == 0 {
		flag.PrintDefaults()
		return
	}

	log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(*path))))
}
