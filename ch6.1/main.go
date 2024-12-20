package main

import (
	"fmt"

)

func main() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	a := make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		a = append(a, Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       100,
		})
	}
	a0 := &a[0]
	fmt.Printf("a0: %p \n", a0)
	fmt.Printf("&a[0]: %p \n", &a[0])
	fmt.Printf("&a[1] = %p \n", &a[1])
	fmt.Println("Done, len(a)=", len(a))
}