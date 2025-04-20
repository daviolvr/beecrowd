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

		if X%2 != 0 {
			X++
		}

		sum := 0

		for i := 0; i < 5; i++ {
			sum += X + i*2
		}

		fmt.Println(sum)
	}
}
