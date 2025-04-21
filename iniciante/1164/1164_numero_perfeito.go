package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var X int
		sumValues := 0
		fmt.Scan(&X)

		for j := 1; j < X; j++ {
			if X%j == 0 {
				sumValues += j
			}
		}

		if X == sumValues {
			fmt.Printf("%d eh perfeito\n", X)
		} else {
			fmt.Printf("%d nao eh perfeito\n", X)
		}
	}
}
