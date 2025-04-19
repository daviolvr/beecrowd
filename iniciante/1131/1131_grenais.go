package main

import "fmt"

func main() {
	vitoriasInter, vitoriasGremio, empates := 0, 0, 0

	for {
		var golsInter, golsGremio int
		fmt.Scan(&golsInter, &golsGremio)

		if golsInter > golsGremio {
			vitoriasInter++
		} else if golsGremio > golsInter {
			vitoriasGremio++
		} else if golsInter == golsGremio {
			empates++
		}

		var command int
		fmt.Println("Novo grenal (1-sim 2-nao)")
		fmt.Scan(&command)

		if command == 1 {
			continue
		} else if command == 2 {
			break
		}
	}

	grenais := vitoriasInter + vitoriasGremio + empates
	fmt.Printf("%d grenais\n", grenais)

	fmt.Printf("Inter:%d\n", vitoriasInter)
	fmt.Printf("Gremio:%d\n", vitoriasGremio)
	fmt.Printf("Empates:%d\n", empates)

	if vitoriasInter > vitoriasGremio {
		fmt.Println("Inter venceu mais")
	} else if vitoriasGremio > vitoriasInter {
		fmt.Println("Gremio venceu mais")
	} else if vitoriasInter == vitoriasGremio {
		fmt.Println("Nao houve vencedor")
	}
}
