package main

import (
	"fmt"
)

func main() {
	var X float64
	fmt.Scan(&X)

	var N [100]float64
	N[0] = X

	fmt.Printf("N[0] = %.4f\n", N[0])

	for i := 1; i < len(N); i++ {
		N[i] = N[i-1] / 2
		fmt.Printf("N[%d] = %.4f\n", i, N[i])
	}
}
