package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)

		if Y == 0 {
			fmt.Println("divisao impossivel")
			continue
		}

		division := float64(X) / float64(Y)
		fmt.Printf("%.1f\n", division)
	}
}
