package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("Cannot Sqrt negative input value:", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z:= x / 2
	z_prev := 0.0
	epsilon := 0.0001
	for i:= 0; math.Abs(z_prev - z) >= epsilon; i++ {
	   z_prev = z
	   z -= (z*z - x) / (2*z)
	   fmt.Println(z, z_prev)
	}
	return z,nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
