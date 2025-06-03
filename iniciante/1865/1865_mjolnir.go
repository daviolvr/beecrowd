package main

import (
	"fmt"
)

func main() {
	var C int
	fmt.Scan(&C)

	for i := 0; i < C; i++ {
		var name string
		var force int
		fmt.Scan(&name, &force)

		if name == "Thor" {
			fmt.Println("Y")
		} else {
			fmt.Println("N")
		}
	}
}
