package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	val2, val1 := 0, 1
	return func() int {
		val := val2
		val2, val1 = val1, val+val1

		return val
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
