package main

import (
	"log"
	"net/http"
)

// TODO: make an env variable
const port string = "8080"

func main() {
	// serves index.html
	// TODO: remove after portal is deployed
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// handles a file upload, returns id of image
	// TODO: cleanup uploads after a certain period of time
	http.HandleFunc("/upload", uploadImage)

	// pass in params to alter an image id, returns an id of the file created, original image is not altered
	// 1) inputs for desired board size (need board sizes) - should size image to board and color each square correctly, save the image to pixelated/[id].png
	http.HandleFunc("/pixelate", pixelateImage)

	// pass in the id to download an image
	http.HandleFunc("/download", downloadPixelArt)

	// TODO: better logging
	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
