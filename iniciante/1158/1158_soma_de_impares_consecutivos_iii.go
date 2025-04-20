package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)

		if X%2 == 0 {
			X++
		}

		sum := 0
		for j := 0; j < Y; j++ {
			sum += X + j*2
		}

		fmt.Println(sum)
	}
}
