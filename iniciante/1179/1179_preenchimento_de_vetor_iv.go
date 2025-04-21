package main

import (
	"fmt"
)

func main() {
	var even [5]int
	var odd [5]int
	countEven := 0
	countOdd := 0

	for i := 0; i < 15; i++ {
		var N int
		fmt.Scan(&N)

		if N%2 == 0 {
			even[countEven] = N
			countEven++

			if countEven == len(even) {
				for j := 0; j < len(even); j++ {
					fmt.Printf("par[%d] = %d\n", j, even[j])
				}
				countEven = 0
			}
		} else {
			odd[countOdd] = N
			countOdd++

			if countOdd == len(odd) {
				for j := 0; j < len(odd); j++ {
					fmt.Printf("impar[%d] = %d\n", j, odd[j])
				}
				countOdd = 0
			}
		}
	}

	for i := 0; i < countOdd; i++ {
		fmt.Printf("impar[%d] = %d\n", i, odd[i])
	}
	for i := 0; i < countEven; i++ {
		fmt.Printf("par[%d] = %d\n", i, even[i])
	}
}
