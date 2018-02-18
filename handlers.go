package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
}

func handleFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
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
		// fmt.Println(handler)

		image, err := os.Open(string(handler.Filename))
		if err != nil {
			os.Exit(1)
		}

		defer image.Close()

		fmt.Println(image)
	}
}
