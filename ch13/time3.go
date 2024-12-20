package main

import (
	"fmt"
	"time"
)

func main() {
	t1, err := time.Parse(time.DateTime, "1957-11-22 05:22:15")
	if err != nil {
		panic(err)
	}
	fmt.Println(t1)
}
