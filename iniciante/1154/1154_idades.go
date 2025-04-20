package main

import (
	"fmt"
)

func main() {
	var N, sum, countValues int

	for {
		fmt.Scan(&N)

		if N < 0 {
			break
		}

		sum += N
		countValues++
	}

	mean := float64(sum) / float64(countValues)
	fmt.Printf("%.2f\n", mean)
}
