package main

import (
	"fmt"
)

func main() {
	// condicao: i + j >= 12
	var operation string
	var matriz [12][12]float64
	var soma float64
	var count int

	fmt.Scan(&operation)
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			if i+j >= 12 {
				soma += matriz[i][j]
				count++
			}
		}
	}

	if operation == "S" {
		fmt.Printf("%.1f\n", soma)
	} else if operation == "M" {
		media := soma / float64(count)
		fmt.Printf("%.1f\n", media)
	}
}
