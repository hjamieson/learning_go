package main

import (
	"compress/gzip"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	out, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Open file error: ", err)
	}
	gzin, err := gzip.NewReader(out)
	buffer := make([]byte, 100)
	count, err := gzin.Read(buffer)
	gzin.Close()
	log.Println("Read", "bytes read", count, "file", fileName)
	log.Println("Read", "data", string(buffer))
}