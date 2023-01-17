package handlers

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to Taocoder Go API");
}