package main

import (
	"fmt"
)

func main() {
	var X, Y int
	fmt.Scan(&X, &Y)

	for i := 1; i <= Y; i += X {
		for j := 0; j < X; j++ {
			if i+j <= Y {
				if j == X-1 || i+j == Y {
					fmt.Printf("%d", i+j)
				} else {
					fmt.Printf("%d ", i+j)
				}
			}
		}
		fmt.Println()
	}
}
