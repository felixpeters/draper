// Package web provides the web server infrastructure for the application.
package web

import (
	"net/http"
	"time"
)

// NewHTTPServer returns a new HTTP server instance.
// The address of the server is set to addr.
// Routes can be configured via router, a standard http.Handlerrouter.
func NewHTTPServer(addr string, router http.Handler) *http.Server {
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}
