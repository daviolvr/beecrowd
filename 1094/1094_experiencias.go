package main

import (
	"fmt"
)

func main() {
	var N int
	countRatos, countSapos, countCoelhos, total := 0, 0, 0, 0
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var num int
		var tipo string
		fmt.Scan(&num, &tipo)

		if tipo == "C" {
			countCoelhos += num
		} else if tipo == "S" {
			countSapos += num
		} else if tipo == "R" {
			countRatos += num
		}

		total += num
	}

	totalFloat := float64(total)
	countCoelhosFloat := float64(countCoelhos)
	countRatosFloat := float64(countRatos)
	countSaposFloat := float64(countSapos)

	percentualCoelhos := float64((countCoelhosFloat / totalFloat) * 100.0)
	percentualRatos := float64((countRatosFloat / totalFloat) * 100.0)
	percentualSapos := float64((countSaposFloat / totalFloat) * 100.0)

	fmt.Printf("Total: %d cobaias\n", total)
	fmt.Printf("Total de coelhos: %d\n", countCoelhos)
	fmt.Printf("Total de ratos: %d\n", countRatos)
	fmt.Printf("Total de sapos: %d\n", countSapos)
	fmt.Printf("Percentual de coelhos: %.2f %%\n", percentualCoelhos)
	fmt.Printf("Percentual de ratos: %.2f %%\n", percentualRatos)
	fmt.Printf("Percentual de sapos: %.2f %%\n", percentualSapos)
}
