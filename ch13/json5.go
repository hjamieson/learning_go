package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

func main() {
	streamData := `
	{"name": "John", "age": 30}
	{"name": "Jane", "age": 25}
	{"name": "Jack", "age": 40}`

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	people := make([]Person, 0, 3)
	var peep Person
	dec := json.NewDecoder(strings.NewReader(streamData))
	for {
		err := dec.Decode(&peep)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		people = append(people, peep)
	}
	fmt.Printf("people: %v \n", people)
}
