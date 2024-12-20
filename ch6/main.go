package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func makePerson(firstName, lastName string, age int) Person {
	return Person{
		firstName,
		lastName,
		age,
	}
}

func makePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		firstName,
		lastName,
		age,
	}
}

func main() {
	fmt.Println(makePerson("Hugh", "Jamieson", 67))
	fmt.Println(makePersonPointer("Hugh", "Jamieson", 67))
}