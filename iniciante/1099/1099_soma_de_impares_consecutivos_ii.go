package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var X, Y, sumOdd int
		fmt.Scan(&X, &Y)

		if X > Y {
			for i := Y + 1; i < X; i++ {
				if i%2 != 0 {
					sumOdd += i
				}
			}
		} else if Y > X {
			for i := X + 1; i < Y; i++ {
				if i%2 != 0 {
					sumOdd += i
				}
			}
		}

		fmt.Println(sumOdd)
	}
}
