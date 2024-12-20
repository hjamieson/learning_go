package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]int{}
	m["hugh"] = 1
	m["kay"] = 2
	m["alex"] = 3
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	/*
		prints: {"alex":3,"hugh":1,"kay":2}
	*/
}
