package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	N := [1000]int{}

	for i := 0; i < 1000; i++ {
		N[i] = i % T
		fmt.Printf("N[%d] = %d\n", i, N[i])
	}
}
