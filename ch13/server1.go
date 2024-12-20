package main

import (
	"net/http"
	"time"
)

func main() {
	person := http.NewServeMux()
	person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Greetings!\n"))
	})
	dog := http.NewServeMux()
	dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Woof!\n"))
	})
	hndlr := http.NewServeMux()
	hndlr.Handle("/person/", http.StripPrefix("/person", person))
	hndlr.Handle("/dog/", http.StripPrefix("/dog", dog))
	s := http.Server{
		Addr:        ":3000",
		Handler:     hndlr,
		ReadTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}
