package main

import "fmt"

func main() {
	var S float64

	for i := 1; i <= 100; i++ {
		S += 1 / float64(i)
	}

	fmt.Printf("%.2f\n", S)
}
