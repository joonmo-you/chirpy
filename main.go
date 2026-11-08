package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	FilePathRoot = "."
	Port         = 8080
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("")))
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", Port),
		Handler: mux,
	}
	log.Printf("Serving files from %s on port: %d\n", FilePathRoot, Port)
	log.Fatal(s.ListenAndServe())
}
