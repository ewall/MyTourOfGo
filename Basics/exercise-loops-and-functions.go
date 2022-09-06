package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, ": ", z)
	}
	return z
}

func main() {
	answer := Sqrt(2)
	fmt.Println("My answer: ", answer)
	fmt.Println("Truth: ", math.Sqrt(2))
	fmt.Println("Difference: ", math.Sqrt(2) - answer)
}
