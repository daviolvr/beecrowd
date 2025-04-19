package main

import (
	"fmt"
)

func main() {
	n1, n2, n3, n4, n5, n6 := 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)
	fmt.Scan(&n4)
	fmt.Scan(&n5)
	fmt.Scan(&n6)

	countPositives := 0
	total := 0.0
	valuesToCompare := []float64{n1, n2, n3, n4, n5, n6}

	for _, num := range valuesToCompare {
		if num < 0 {
			continue
		}

		countPositives++
		total += num
	}

	media := total / float64(countPositives)

	fmt.Printf("%d valores positivos\n", countPositives)
	fmt.Printf("%.1f\n", media)

}
