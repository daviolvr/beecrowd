package main

import (
	"fmt"
)

func main() {
	var A [100]float64
	for i := 0; i < len(A); i++ {
		fmt.Scan(&A[i])
		if A[i] <= 10 {
			fmt.Printf("A[%d] = %.1f\n", i, A[i])
		}
	}
}
