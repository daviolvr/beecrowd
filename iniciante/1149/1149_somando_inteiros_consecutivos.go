package main

import (
	"fmt"
)

func main() {
	var A, N, sum int
	fmt.Scan(&A, &N)

	for N <= 0 {
		fmt.Scan(&N)
	}

	for i := 0; i <= N-1; i++ {
		sum += A + i
	}

	fmt.Printf("%d\n", sum)
}
