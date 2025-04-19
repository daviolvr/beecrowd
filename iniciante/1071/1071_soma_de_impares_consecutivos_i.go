package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)

	sum := 0

	for i := y + 1; i < x; i++ {
		if i%2 != 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
