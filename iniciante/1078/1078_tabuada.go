package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", i, N, i*N)
	}
}
