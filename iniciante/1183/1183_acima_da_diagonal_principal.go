package main

import (
	"fmt"
)

func main() {
	var matriz [12][12]float64
	var operation string
	var soma float64
	var countDiagPrincipal int

	fmt.Scan(&operation)

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			if i < j {
				soma += matriz[i][j]
				countDiagPrincipal++
			}
		}
	}

	if operation == "S" {
		fmt.Printf("%.1f\n", soma)
	} else if operation == "M" {
		media := soma / float64(countDiagPrincipal)
		fmt.Printf("%.1f\n", media)
	}
}
