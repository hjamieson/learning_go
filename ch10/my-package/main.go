// This is the main file that will use the functions from the other packages	
package main

import (
	"fmt"
	"my-package/do-format"
	"my-package/math"
)

// main is the entry point for the application
func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}