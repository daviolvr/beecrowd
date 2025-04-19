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

	evenCount := 0
	valuesToCheck := []int{n1, n2, n3, n4, n5}

	for _, num := range valuesToCheck {
		if num%2 == 0 {
			evenCount++
		}
	}

	fmt.Printf("%d valores pares\n", evenCount)
}
