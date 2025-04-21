package main

import (
	"fmt"
)

// Here i used range only once, but i could have used it on the first loop as well.
// I chose not to use it in the first loop just for practice purposes.
func main() {
	var N [20]int

	for i := 0; i < len(N); i++ {
		fmt.Scan(&N[i])
	}

	for i := 0; i < len(N)/2; i++ {
		j := len(N) - 1 - i
		N[j], N[i] = N[i], N[j]
	}

	for i, v := range N {
		fmt.Printf("N[%d] = %d\n", i, v)
	}
}
