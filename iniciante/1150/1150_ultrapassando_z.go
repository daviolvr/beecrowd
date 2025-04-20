package main

import (
	"fmt"
)

func main() {
	var X, Z, countValues, sum int
	fmt.Scan(&X)
	fmt.Scan(&Z)

	for Z <= X {
		fmt.Scan(&Z)
	}

	for i := X; i < Z; i++ {
		sum += i
		countValues++
		if sum > Z {
			break
		}
	}

	fmt.Println(countValues)
}
