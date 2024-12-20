package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	bits, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bits))
	var o Order
	err = json.Unmarshal(bits, &o)
	if err != nil {
		panic(err)
	}
	fmt.Println(o)

	o2 := Order{
		ID:          "022",
		DateOrdered: time.Now(),
		CustomerID:  "14",
		Items: []Item{
			Item{
				ID:   "A1",
				Name: "Shirt",
			},
			Item{
				ID:   "A2",
				Name: "Pantz",
			},
		},
	}
	fmt.Println("o2: ", o2)

	o2Bits, err := json.Marshal(&o2)
	fmt.Println("o2 JSON: ", string(o2Bits))
}
