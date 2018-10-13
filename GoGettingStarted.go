package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	for z := 1.0; ; {
		prev := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z, "|", prev)
		if float32(prev) == float32(z) {
			return z
		}
	}
}

func main() {
	n := 2.
	fmt.Println("state: ", Sqrt(n) == math.Sqrt(n))
}
