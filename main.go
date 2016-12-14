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

type RequestLogger interface {
	Log(method string, url string, duration int, code int, size int)
}

type loggingHandler struct {
	logger  log.Logger
	handler http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t1 := time.Now()
	logrw := NewLoggingResponseWriter(w)
	h.handler.ServeHTTP(logrw, req)
	t2 := time.Now()
	h.logger.Log("method", req.Method, "url", req.URL.String(), "responseTime", t2.Sub(t1), "code", logrw.Status(), "size", logrw.Size())
}

func NewLoggingResponseWriter(w http.ResponseWriter) loggingResponseWriter {
	var logger loggingResponseWriter = &responseLogger{w: w}
	return logger
}

type loggingResponseWriter interface {
	http.ResponseWriter
	http.Flusher
	Status() int
	Size() int
}

type responseLogger struct {
	w      http.ResponseWriter
	status int
	size   int
}

func (l *responseLogger) Header() http.Header {
	return l.w.Header()
}

func (l *responseLogger) Write(b []byte) (int, error) {
	if l.status == 0 {
		// The status will be StatusOK if WriteHeader has not been called yet
		l.status = http.StatusOK
	}
	size, err := l.w.Write(b)
	l.size += size
	return size, err
}

func (l *responseLogger) WriteHeader(s int) {
	l.w.WriteHeader(s)
	l.status = s
}

func (l *responseLogger) Status() int {
	return l.status
}

func (l *responseLogger) Size() int {
	return l.size
}

func (l *responseLogger) Flush() {
	f, ok := l.w.(http.Flusher)
	if ok {
		f.Flush()
	}
}

var logger log.Logger

func init() {

	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.NewContext(logger).With("ts", log.DefaultTimestamp, "caller", log.DefaultCaller)
}

func LoggingHandler(l log.Logger, h http.Handler) http.Handler {
	return loggingHandler{l, h}
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
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello World")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HelloWorld)
	r.HandleFunc("/json", JsonHelloWorld)
	loggedRouter := LoggingHandler(logger, recoverHandler(r))

	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         os.Getenv("HTTP_ADDR"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Log("info", "HTTP server started on port 8080")
	srv.ListenAndServe()
}
