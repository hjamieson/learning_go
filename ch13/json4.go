package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Book struct {
	id     int
	Title  string
	Author string
}

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/boom", boom)
	http.HandleFunc("/book", getBook)
	http.HandleFunc("/dummy", getDummy)
	log.Println("starting server...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/index.html", http.StatusFound)
}

func boom(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "dumb", http.StatusBadRequest)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	book := Book{Title: "The Firm", Author: "Jon Grisham"}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func getDummy(w http.ResponseWriter, r *http.Request) {
	book := Book{Title: "Slaughterhouse Five", Author: "Kurt Vonnegut"}

	var dummy = template.Must(template.New("dummy").Parse(`
		<html>
			<head>
			</head>
			<body>
			<h1>{{.Title}}</h1>
			<h3>{{.Author}}</h3>
			<p>This is a good book</p>
			</body>
			</html>
			`))
	dummy.Execute(w, book)
}
