package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	j := 1
	for i := 1; i <= N; i++ {
		fmt.Printf("%d %d %d PUM\n", j, j+1, j+2)
		j += 4
	}
}
