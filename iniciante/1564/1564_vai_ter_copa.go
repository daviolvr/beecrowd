package main

import (
	"fmt"
)

func main() {
	var N int
	for {
		_, err := fmt.Scan(&N)
		if err != nil {
			break
		}

		if N == 0 {
			fmt.Println("vai ter copa!")
		} else {
			fmt.Println("vai ter duas!")
		}
	}
}
