package main

import (
	"fmt"
)

func main() {
	for {
		var X, Y int
		fmt.Scan(&X, &Y)

		if X == 0 || Y == 0 {
			break
		}

		if X > 0 && Y > 0 {
			fmt.Println("primeiro")
		}

		if X < 0 && Y > 0 {
			fmt.Println("segundo")
		}

		if X < 0 && Y < 0 {
			fmt.Println("terceiro")
		}

		if X > 0 && Y < 0 {
			fmt.Println("quarto")
		}
	}
}
