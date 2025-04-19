package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var n1, n2, n3 float64
		fmt.Scan(&n1, &n2, &n3)

		weighted_average := (n1*2 + n2*3 + n3*5) / 10
		fmt.Printf("%.1f\n", weighted_average)
	}
}
