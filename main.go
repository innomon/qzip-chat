package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:public
var content embed.FS

func main() {
	port := flag.String("port", "7070", "Port to serve on")
	flag.Parse()

	// Use fs.Sub to create a sub-filesystem rooted at the "static" directory.
	// This ensures that http.FileServer can correctly find files relative to the base directory.
	subFS, err := fs.Sub(content, "public")
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.FS(subFS))
	http.Handle("/", fs)

	log.Printf("Serving on port %s", *port)
	err = http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
