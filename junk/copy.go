package main

import (
	"errors"
	"fmt"
)

func main() {
	e1 := errors.New("1st error")
	e2 := errors.New("2nd error")
	fmt.Println("e1 == e1", e1 == e1)
	fmt.Println("e2 == e2", e2 == e2)
	fmt.Println("e1 == e2", e1 == e2)
	fmt.Println("e1 == errors.new(1st error)", e1 == errors.New("1st error"))
	e3 := errors.New("1st error")
	fmt.Println("e1 == e3", e1 == e3)
	fmt.Println("errors.Is(e1, e3)", errors.Is(e1, e3))
	fmt.Println("errors.Is(e1, e1)", errors.Is(e1, e1))
}
