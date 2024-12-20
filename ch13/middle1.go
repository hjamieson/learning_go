package main

import (
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)
	wrappedMux := AddHeader(AnotherHeader(mux))
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: wrappedMux,
	}

	server.ListenAndServe()
}

func AddHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Hugh", "was here")
		h.ServeHTTP(w, r)
		slog.Info("AddHeader processsed request")
	})
}

func AnotherHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Another", "was here")
		h.ServeHTTP(w, r)
		slog.Info("AddHeader processsed request")
	})
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
	slog.Info("Hello processed request")
}
