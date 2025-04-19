package main

import "fmt"

func main() {
	i := 1
	for i <= 9 {
		for j := 7; j >= 5; j-- {
			fmt.Printf("I=%d J=%d\n", i, j)
		}

		i += 2
	}
}
