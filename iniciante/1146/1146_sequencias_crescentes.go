package main

import (
	"fmt"
)

func main() {
	for {
		var X int
		fmt.Scan(&X)

		if X == 0 {
			break
		}

		for i := 1; i <= X; i++ {
			if i == X {
				fmt.Printf("%d", i)
			} else {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Println()
	}
}
