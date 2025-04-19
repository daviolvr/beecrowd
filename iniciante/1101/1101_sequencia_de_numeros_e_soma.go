package main

import (
	"fmt"
)

func main() {
	for {
		var M, N, sum int
		fmt.Scan(&M, &N)

		if M <= 0 || N <= 0 {
			return
		}

		if M > N {
			for i := N; i <= M; i++ {
				fmt.Printf("%d ", i)
				sum += i
			}
		} else {
			for i := M; i <= N; i++ {
				fmt.Printf("%d ", i)
				sum += i
			}
		}

		fmt.Printf("Sum=%d\n", sum)
	}
}
