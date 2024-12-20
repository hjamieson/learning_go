package main

import (
	"fmt"
	"math/rand"
)

func arrays() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}

	// arrays
	var x1 [3]int = [3]int{1, 2, 3}
	fmt.Println("x1[2] = ", x1[2])

	// you can use index expression in array literal to set
	// a specific element:
	var a1 = [5]int{2: 11}
	fmt.Println(a1[2])

	// let compiler assign the length:
	var a2 = [...]uint64{10, 20, 30}
	fmt.Println(a2[1])

	// [...] makes an array: [] makes a slice.

	var a3 []int
	a3 = append(a3, rand.Intn(10))
	a3 = append(a3, rand.Intn(10))
	a3 = append(a3, 42)
	fmt.Println("a3 = ", a3)

	// you can't use nil to initialize.
	a4 := append([]int{}, 22)
	fmt.Println("a4 from nil = ", a4)

	a4 = append(a4, 5, 6, 7, len(a4))
	fmt.Println("a4 now ", a4)

	// varadic input parameters & slices
	a5 := []int{0, 1, 2}
	a4 = append(a4, a5...)
	fmt.Println("a4 = ", a4)

	// length and capacity
	fmt.Printf("len(a4)=%d, cap(a4)=%d \n", len(a4), cap(a4))

	// make creates slices with predef type/capacity
	a6 := make([]int, 4) // all zeros
	a6[0] = 12
	a7 := make([]int, 4, 8) // initial size, capacity
	a7[0] = 12

	// clear zeros out all elements of a slice
	clear(a4)
	fmt.Println("a4 after clear: ", a4)

	// slicing slices:
	// taking a slice of a slice is by-reference; data is not copied. Any
	// change to the slice will be seen all other users of the slice.
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e
	fmt.Println("inside addF(), ", list)
	return list
}

func foo(ll []string) {
	fmt.Println("foo got ", ll)
}
