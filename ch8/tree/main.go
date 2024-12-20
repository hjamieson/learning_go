package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	listDemo()
}

func listDemo() {
	list1 := NewList[int]()
	list1.Add(10)
	list1.Add(20)
	list1.Add(30)
	list1.Add(40)
	list1.Add(50)
	list1.Dump(os.Stdout)
	fmt.Println("Index of 30:", list1.Index(30))
	list1.Insert(25, 2)
	list1.Dump(os.Stdout)
}

func treeDemo() {
	treeCompareInt := func(a, b int) int {
		return a - b
	}

	stringCompare := func(a, b string) int {
		return strings.Compare(a, b)
	}

	tree1 := NewTree(treeCompareInt)
	fmt.Printf("%v\n", tree1)
	tree1.Add(10)
	fmt.Printf("%v\n", tree1)
	tree1.Add(5)
	tree1.Add(15)
	tree1.Add(20)
	tree1.Add(1)

	tree1.Dump(os.Stdout)
	fmt.Println("Contains 5:", tree1.Contains(5))

	tree2 := NewTree[string](stringCompare)
	tree2.Add("hello")
	tree2.Add("goodbye")
	tree2.Add("I")
	tree2.Add("ate")
	tree2.Add("peanuts")
	tree2.Dump(os.Stdout)
	fmt.Println("Contains 'goodbye':", tree2.Contains("goodbye"))

	tree3 := NewTree[Person](Person.orderByName)
	tree3.Add(Person{"John", 21})
	tree3.Add(Person{"Mary", 25})
	tree3.Add(Person{"Abel", 7})
	tree3.Add(Person{"Titan", 3})
	tree3.Add(Person{"Kyle", 33})
	tree3.Dump(os.Stdout)
}
