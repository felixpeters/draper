// Package routes provides all API routes according to the REST principles.
// The routes are assembled in a router implementing the standard http.Handler interface.
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter returns a new router instance.
// All specified routes for the application will be contained.
func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorld)
	r.HandleFunc("/json", jsonHelloWorld)
	return r
}
