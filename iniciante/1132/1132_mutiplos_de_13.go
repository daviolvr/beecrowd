package main

import (
	"fmt"
)

func main() {
	var X, Y, sum int
	fmt.Scan(&X, &Y)

	if X > Y {
		X, Y = Y, X
	}

	for i := X; i <= Y; i++ {
		if i%13 == 0 {
			continue
		}
		sum += i
	}

	fmt.Println(sum)
}
