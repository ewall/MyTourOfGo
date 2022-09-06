package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last, next := 0, 0
	return func() int {
		if last == 0 && next == 0 {
			next += 1
		} else {
			last, next = next, last+next
		}
		return last
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
