package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		if num%2 == 0 {
			fmt.Println("0")
		} else {
			fmt.Println("1")
		}
	}
}
