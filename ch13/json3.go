package main

import (
	"encoding/json"
	"os"
)

type Book struct {
	Title  string
	Author string `json:"contributer,omitempty"`
}

func main() {
	books := make([]Book, 3)
	books = append(books, Book{"The Martian", "Andy Weir"})
	books = append(books, Book{"The Firm", "Grisham"})
	books = append(books, Book{Title: "One Flew Over The Cookoos Nest"})

	tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	//defer os.Remove(tmpFile.Name())
	err = json.NewEncoder(tmpFile).Encode(books)
	if err != nil {
		panic(err)
	}
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}
}
