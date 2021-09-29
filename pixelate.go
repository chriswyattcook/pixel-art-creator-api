package main

import (
	"fmt"
	"net/http"
)

func pixelateImage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pixelate")
}
