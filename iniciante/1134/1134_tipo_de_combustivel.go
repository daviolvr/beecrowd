package main

import (
	"fmt"
)

func main() {
	countAlcool, countGasolina, countDiesel := 0, 0, 0

	for {
		var codigo int
		fmt.Scan(&codigo)

		if codigo == 4 {
			break
		}

		if codigo < 1 || codigo > 4 {
			continue
		}

		if codigo == 1 {
			countAlcool++
		} else if codigo == 2 {
			countGasolina++
		} else if codigo == 3 {
			countDiesel++
		}
	}

	fmt.Println("MUITO OBRIGADO")
	fmt.Printf("Alcool: %d\n", countAlcool)
	fmt.Printf("Gasolina: %d\n", countGasolina)
	fmt.Printf("Diesel: %d\n", countDiesel)
}
