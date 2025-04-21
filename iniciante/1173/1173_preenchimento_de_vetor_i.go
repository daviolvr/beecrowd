package main

import (
	"fmt"
)

func main() {
	var V int
	var N [10]int
	fmt.Scan(&V)

	for i := 0; i < len(N); i++ {
		N[i] = V
		fmt.Printf("N[%d] = %d\n", i, N[i])
		V *= 2
	}
}
