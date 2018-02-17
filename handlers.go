package main

import (
	"net/http"
)

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

}
