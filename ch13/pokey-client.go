package main

import (
	"io"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 6 * time.Second,
	}

	req, err := http.NewRequest("GET", "http://localhost:3000/pokey", nil)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	duration := time.Since(start)
	text, err := io.ReadAll(res.Body)
	slog.Info("response", "body", string(text), "duration", duration)

}
