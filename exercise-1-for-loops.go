package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z:= x / 2
	z_prev := 0.0
	epsilon := 0.0001
	for i:= 0; math.Abs(z_prev - z) >= epsilon; i++ {
	   z_prev = z
	   z -= (z*z - x) / (2*z)
	   fmt.Println(z, z_prev)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
