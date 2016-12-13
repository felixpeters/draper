package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	log "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

type SimpleResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var logger log.Logger

func init() {

	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.NewContext(logger).With("ts", log.DefaultTimestamp, "caller", log.DefaultCaller)
}

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		logger.Log("method", r.Method, "url", r.URL.String(), "responseTime", t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log("panic", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
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

	r := mux.NewRouter()
	r.HandleFunc("/", HelloWorld)
	r.HandleFunc("/json", JsonHelloWorld)
	loggedRouter := loggingHandler(recoverHandler(r))

	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         os.Getenv("HTTP_ADDR"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Log("info", "HTTP server started on port 8080")
	srv.ListenAndServe()
}
