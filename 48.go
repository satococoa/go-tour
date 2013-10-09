package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := x
	for i := 0; i < 20; i++ {
		z = z - ((cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2)))
	}
	return z
}

func main() {
	fmt.Println(math.Pow(2, 1.0/3.0))
	fmt.Println(Cbrt(2))
}
