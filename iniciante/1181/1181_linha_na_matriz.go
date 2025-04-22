package main

import (
	"fmt"
)

func main() {
	var L int
	var T string
	var sum float64

	fmt.Scan(&L)
	fmt.Scan(&T)

	matriz := make([][]float64, 12)
	for i := 0; i < 12; i++ {
		matriz[i] = make([]float64, 12)
		for j := 0; j < 12; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	for j := 0; j < 12; j++ {
		sum += matriz[L][j]
	}

	if T == "S" {
		fmt.Printf("%.1f\n", sum)
	} else if T == "M" {
		media := sum / 12.0
		fmt.Printf("%.1f\n", media)
	}
}
