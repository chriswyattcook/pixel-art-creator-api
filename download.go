package main

import (
	"fmt"
	"net/http"
)

func downloadPixelArt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "download")
}
