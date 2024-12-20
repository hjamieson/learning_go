package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Orderx struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string
}

type Book struct {
	Title  string
	Author string `json:"contributer,omitempty"`
}

func main() {
	o1 := Orderx{ID: "42", DateOrdered: time.Now(), CustomerID: "Frodo"}
	o1Bits, _ := json.Marshal(&o1)
	fmt.Println("o1Bits =", string(o1Bits))

	b1 := Book{"The Martian", "Andy Weir"}
	fmt.Println(b1)
	b1Bytes, _ := json.Marshal(&b1)
	fmt.Println(string(b1Bytes))

	b2, _ := json.Marshal(Book{Title: "The Firm"})
	fmt.Println(string(b2))
	b3, _ := json.Marshal(Book{"The Firm", "Jon Grisham"})
	fmt.Println(string(b3))
}
