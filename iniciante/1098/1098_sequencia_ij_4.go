package main

import "fmt"

func main() {
	for i := 0; i <= 20; i += 2 {
		for j := 1; j <= 3; j++ {
			ii := float64(i) / 10
			jj := ii + float64(j)

			if i%10 == 0 {
				fmt.Printf("I=%.0f J=%.0f\n", ii, jj)
			} else {
				fmt.Printf("I=%.1f J=%.1f\n", ii, jj)
			}
		}
	}
}
