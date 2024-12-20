package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("1hour =", time.Hour)
	fmt.Printf("1hour : %v \n", time.Hour)
	fmt.Printf("1hour : %v \n", int64(time.Hour))
}
