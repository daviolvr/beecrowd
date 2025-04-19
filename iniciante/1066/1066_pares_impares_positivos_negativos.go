package main

import (
	"fmt"
)

func main() {
	n1, n2, n3, n4, n5 := 0, 0, 0, 0, 0
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)
	fmt.Scan(&n4)
	fmt.Scan(&n5)

	values := []int{n1, n2, n3, n4, n5}
	countEven, countOdd, countPositive, countNegative := 0, 0, 0, 0

	for _, num := range values {
		if num%2 == 0 {
			countEven++
		} else {
			countOdd++
		}

		if num > 0 {
			countPositive++
		} else if num < 0 {
			countNegative++
		}
	}

	fmt.Printf("%d valor(es) par(es)\n", countEven)
	fmt.Printf("%d valor(es) impar(es)\n", countOdd)
	fmt.Printf("%d valor(es) positivo(s)\n", countPositive)
	fmt.Printf("%d valor(es) negativo(s)\n", countNegative)
}
