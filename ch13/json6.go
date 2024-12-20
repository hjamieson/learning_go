package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	type Person struct {
		Name     string    `json:"name"`
		Birthday time.Time `json:"dob"`
	}

	hugh := Person{"Hugh", time.Now()}
	bytes, err := json.Marshal(hugh)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	/*
		prints: {"name":"Hugh","dob":"2024-11-21T16:36:45.461757-05:00"}
	*/
}
