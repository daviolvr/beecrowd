package main

import (
	"fmt"
)

func main() {
	var greater int
	var greaterPos int
	for i := 0; i < 100; i++ {
		var n int
		fmt.Scan(&n)

		if n > greater {
			greater = n
			greaterPos = i + 1
		}
	}

	fmt.Println(greater)
	fmt.Println(greaterPos)
}
