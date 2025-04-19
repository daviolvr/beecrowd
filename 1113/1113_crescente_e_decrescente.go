package main

import (
	"fmt"
)

func main() {
	for {
		var X, Y int
		fmt.Scan(&X, &Y)

		if X == Y {
			break
		}

		if X > Y {
			fmt.Println("Decrescente")
			continue
		}

		fmt.Println("Crescente")
	}
}
