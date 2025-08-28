package main

import (
	"fmt"
)

func main() {
	var col int
	var operation string
	var resultado float64
	var matriz [12][12]float64

	fmt.Scan(&col)
	fmt.Scan(&operation)

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	for i := 0; i < 12; i++ {
		resultado += matriz[i][col]
	}

	if operation == "M" {
		resultado /= 12.0
	}

	fmt.Printf("%.1f\n", resultado)
}
