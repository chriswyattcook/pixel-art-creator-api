package main

import (
	"log"
	"net/http"
)

const serverAddress string = ":8080"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/upload", uploadFile)

	log.Printf("Listening on %s", serverAddress)

	log.Fatal(http.ListenAndServe(serverAddress, nil))
}
