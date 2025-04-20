package main

import "fmt"

func main() {
	var S float64
	j := 1

	for i := 1; i <= 39; i += 2 {
		S += float64(i) / float64(j)
		j *= 2
	}

	fmt.Printf("%.2f\n", S)
}
