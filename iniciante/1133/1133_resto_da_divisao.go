package main

import (
	"fmt"
)

func main() {
	var X, Y int
	fmt.Scan(&X)
	fmt.Scan(&Y)

	if X > Y {
		X, Y = Y, X
	}

	for i := X + 1; i < Y; i++ {
		if i%5 == 2 || i%5 == 3 {
			fmt.Println(i)
		}
	}
}
