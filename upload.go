package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

func uploadImage(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Fprintf(w, "error: %s", err)
	}
	defer file.Close()

	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}

	tempFile, err := ioutil.TempFile("upload", uuid.String()+"-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	imageId := strings.Split(strings.Split(tempFile.Name(), "/")[len(strings.Split(tempFile.Name(), "/"))-1], ".")[0]

	fmt.Fprintf(w, "created: %s", imageId)
}
