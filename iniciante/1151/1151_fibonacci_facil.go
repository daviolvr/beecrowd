package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	if N == 1 {
		fmt.Println(0)
		return
	}

	n1, n2 := 0, 1
	fmt.Printf("%d %d", n1, n2)

	for i := 2; i < N; i++ {
		n3 := n1 + n2
		fmt.Printf(" %d", n3)
		n1 = n2
		n2 = n3
	}
	fmt.Println()
}
