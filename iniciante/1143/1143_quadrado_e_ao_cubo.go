package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	// You can also use math.Pow from "math"
	for i := 1; i <= N; i++ {
		squared := i * i
		cubed := i * i * i
		fmt.Printf("%d %d %d\n", i, squared, cubed)
	}
}
