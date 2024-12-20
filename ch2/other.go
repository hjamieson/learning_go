package main

import (
	"fmt"
	"math"
)

const x int64 = 10

const (
	idKey = "id"
	nameKey = "name"
)
const z = 20 * 10

func other(){
	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)

}

func ex1(){
	var i int = 20
	var f float64 = float64(i)
	fmt.Printf("i is %d, f is %f \n", i,f)
}

func ex2() {
	const value = 42
	var i int = value
	var f float64 = value
	fmt.Printf("i is %d, f is %f \n", i,f)
}

func ex3(){
	var (
		b byte
		smallI int
		bigI uint64
	)
	b = math.MaxUint8
	smallI = math.MaxInt
	bigI = math.MaxUint64
	fmt.Printf("b is %d, smallI is %d, bigI is %d \n", b, smallI, bigI)
	b +=1
	smallI +=1
	bigI +=1
	fmt.Printf("b is %d, smallI is %d, bigI is %d \n", b, smallI, bigI)


}