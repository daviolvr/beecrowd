package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		var X int
		fmt.Scan(&X)
		sumDividers := 0

		for j := 1; j <= X; j++ {
			if X%j == 0 {
				sumDividers++
			}
		}

		if sumDividers == 2 {
			fmt.Printf("%d eh primo\n", X)
		} else {
			fmt.Printf("%d nao eh primo\n", X)
		}
	}
}
