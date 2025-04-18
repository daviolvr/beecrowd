package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Scan(&x)

	if x%2 == 0 {
		x++
	}

	for i := 0; i < 6; i++ {
		fmt.Println(x + i*2)
	}
}
