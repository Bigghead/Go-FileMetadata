package main

import (
	"fmt"
	"net/http"
)

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
}

func handleFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

}
