package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var N int
		fmt.Scan(&N)

		n1, n2 := 0, 1
		var fib int

		if N == 0 {
			fib = 0
		} else if N == 1 {
			fib = 1
		} else {
			for j := 2; j <= N; j++ {
				fib = n1 + n2
				n1 = n2
				n2 = fib
			}
			fib = n2
		}

		fmt.Printf("Fib(%d) = %d\n", N, fib)
	}
}
