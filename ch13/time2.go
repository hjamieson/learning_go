package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now()
	fmt.Println("sleeping")
	time.Sleep(3 * time.Second)
	et := time.Since(st)

	fmt.Println("Awake", et)
	fmt.Println("Awake(sec)", et.Seconds())
	fmt.Println("Awake(sec)", et.Round(time.Second))
}
