package main

import "fmt"

func main() {
	i := 1
	jValue := 7
	for i <= 9 {
		for j := jValue; j >= jValue-2; j-- {
			fmt.Printf("I=%d J=%d\n", i, j)
		}

		i += 2
		jValue += 2
	}
}
