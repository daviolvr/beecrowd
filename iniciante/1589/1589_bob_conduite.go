package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var R1, R2 int
		fmt.Scan(&R1, &R2)
		minPossible := R1 + R2
		fmt.Println(minPossible)
	}
}
