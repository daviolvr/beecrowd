package main

import (
	"fmt"
)

func main() {
	var N int
	fatorialN := 1
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		fatorialN *= i
	}

	fmt.Println(fatorialN)
}
