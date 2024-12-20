package main

import "strings"

type Person struct {
	name string
	age  int
}
func (p Person) orderByName(other Person) int {
	return strings.Compare(p.name, other.name)
}
