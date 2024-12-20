package main

import (
	"log/slog"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/pokey", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request")
		time.Sleep(5 * time.Second)
		w.Write([]byte("Slowpoke says: 'Hi there!'"))
	})
	http.ListenAndServe(":3000", nil)
}
