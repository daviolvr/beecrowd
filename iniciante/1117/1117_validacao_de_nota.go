package main

import (
	"fmt"
)

func main() {
	validCount := 0
	sum := 0.0

	for {
		var nota float64
		fmt.Scan(&nota)

		if nota < 0 || nota > 10 {
			fmt.Println("nota invalida")
			continue
		}

		sum += nota
		validCount++

		if validCount == 2 {
			mean := sum / 2
			fmt.Printf("media = %.2f\n", mean)
			break
		}
	}
}
