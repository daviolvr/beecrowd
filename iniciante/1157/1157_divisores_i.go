package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		if N%i == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}
