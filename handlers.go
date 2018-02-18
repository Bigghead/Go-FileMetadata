package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type metaData struct {
	Name  string
	Size  int64
	Time  time.Time
	IsDir bool
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
}

func handleFile(w http.ResponseWriter, r *http.Request) {

	// ==== read file upload, make a new file ===== //
	// ==== copy contents into a new file, read new file metadata ===== //

	if r.Method == "POST" {
		r.ParseMultipartForm(24)
		file, handler, err := r.FormFile("uploadFile")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer f.Close()

		io.Copy(f, file)

		image, err := os.Stat(string(handler.Filename))
		if err != nil {
			os.Exit(1)
		}

		fileMetadata := metaData{
			Name:  image.Name(),
			Size:  image.Size(),
			Time:  image.ModTime(),
			IsDir: image.IsDir(),
		}

		a, err := json.Marshal(fileMetadata)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(fileMetadata)
	}
}
