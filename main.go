package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type SimpleResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func JsonHelloWorld(w http.ResponseWriter, r *http.Request) {
	response := SimpleResponse{true, "Hello World"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/json", JsonHelloWorld)
	http.ListenAndServe(":8080", nil)
}
