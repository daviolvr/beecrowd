package main

import (
	"fmt"
)

func main() {
	var operation byte
	fmt.Scanf("%c\n", &operation)

	var matrix [12][12]float64
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	var sum float64
	var count int

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			if i > j && i+j >= 12 {
				sum += matrix[i][j]
				count++
			}
		}
	}

	if operation == 'M' {
		average := sum / float64(count)
		fmt.Printf("%.1f\n", average)
	} else {
		fmt.Printf("%.1f\n", sum)
	}
}
