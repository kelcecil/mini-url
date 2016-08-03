package main

import (
	"log"
	"net/http"
)

// main .. This just initializes the storage, hands it off to the main http
// handler, and starts the server.
func main() {
	log.Print("Initializing storage")
	storage := InitMapStorage()

	http.Handle("/", ShortUrlForwardingHandler{
		Storage: storage,
	})

	log.Print("Starting server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server; %v", err.Error())
	}
}
