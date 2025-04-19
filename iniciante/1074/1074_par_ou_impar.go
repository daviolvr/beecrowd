package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var X int
		fmt.Scan(&X)

		if X == 0 {
			fmt.Println("NULL")
			continue
		}

		if X%2 == 0 {
			if X > 0 {
				fmt.Println("EVEN POSITIVE")
			}

			if X < 0 {
				fmt.Println("EVEN NEGATIVE")
			}
		}

		if X%2 != 0 {
			if X > 0 {
				fmt.Println("ODD POSITIVE")
			}

			if X < 0 {
				fmt.Println("ODD NEGATIVE")
			}
		}
	}
}
