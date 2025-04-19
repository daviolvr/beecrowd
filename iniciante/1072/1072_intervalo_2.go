package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	withinTheRange, outOfRange := 0, 0

	for i := 0; i < N; i++ {
		var X int
		fmt.Scan(&X)

		if X >= 10 && X <= 20 {
			withinTheRange++
			continue
		}

		outOfRange++
	}

	fmt.Printf("%d in\n", withinTheRange)
	fmt.Printf("%d out\n", outOfRange)
}
