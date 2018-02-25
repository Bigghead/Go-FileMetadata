package main

import (
	"log"
	"net/http"
)

func main() {

	// http.HandleFunc("/", homeRoute)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.HandleFunc("/file", handleFile)
	log.Fatal(http.ListenAndServe(":9000", nil))

}
