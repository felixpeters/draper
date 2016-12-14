package routes

import (
	"encoding/json"
	"io"
	"net/http"
)

// simpleResponse delivers success status and a message
type simpleResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// helloWorld is a handler that responds with the string "Hello World"
func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

// jsonHelloWorld is a handler that responds with JSON-encoded "Hello World"
func jsonHelloWorld(w http.ResponseWriter, r *http.Request) {
	sr := simpleResponse{true, "Hello World"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sr)
}
