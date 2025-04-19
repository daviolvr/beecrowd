package main

import (
	"fmt"
)

func main() {
	for {
		validCount := 0
		sum := 0.0

		for validCount < 2 {
			var nota float64
			fmt.Scan(&nota)

			if nota < 0 || nota > 10 {
				fmt.Println("nota invalida")
				continue
			}

			sum += nota
			validCount++
		}

		mean := sum / 2
		fmt.Printf("media = %.2f\n", mean)

		for {
			fmt.Println("novo calculo (1-sim 2-nao)")
			var command int
			fmt.Scan(&command)

			if command == 1 {
				break
			} else if command == 2 {
				return
			}
		}
	}
}
