package main

import (
	"compress/gzip"
	"log"
	"log/slog"
	"os"
)

func main() {
	// Write a gzip file
	fileName := "test.txt.gz"
	static_text := "Now is the time for all good ℷ to come to the ℏ of ∞."
	out, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Create file error: ", err)
	}
	gzout := gzip.NewWriter(out)
	count, err := gzout.Write([]byte(static_text))
	if err != nil {
		log.Fatal("Write error: ", err)
	}
	slog.Info("Write", "bytes written",count, "file", fileName)
	gzout.Flush()
	gzout.Close()
	out.Close()

}
