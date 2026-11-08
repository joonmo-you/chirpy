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
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(""))))
	mux.HandleFunc("/healthz", handlerReadiness)
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", Port),
		Handler: mux,
	}
	log.Printf("Serving files from %s on port: %d\n", FilePathRoot, Port)
	log.Fatal(s.ListenAndServe())
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
