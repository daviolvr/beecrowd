package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		if i == N {
			fmt.Print("Ho!\n")
			break
		}

		fmt.Print("Ho ")
	}
}
