package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	Tolerance := math.Pow(10.0, -8.0)
	z := 1.0
	for {
		diff := z*z - x
		z -= diff / (2 * x)
		if math.Abs(diff) < Tolerance {
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
