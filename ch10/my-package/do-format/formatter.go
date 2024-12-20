// Package format provides formatting functions for numbers
// 
// It might work; it might not.
package format

import "fmt"

// Number returns a string formatted number
// 
// more info at [Whoopie ky yay mf].
// 
// [Whoopie ky yay mf]: https://www.youtube.com/watch?v=7KXGZAEWvV0
func Number(num int) string {
	return fmt.Sprintf("The number is %d", num)
}