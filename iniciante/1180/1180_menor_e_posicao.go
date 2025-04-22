package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	X := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&X[i])
	}

	menor := X[0]
	posicao := 0

	for i, v := range X {
		if v < menor {
			menor = v
			posicao = i
		}
	}

	fmt.Printf("Menor valor: %d\n", menor)
	fmt.Printf("Posicao: %d\n", posicao)
}
