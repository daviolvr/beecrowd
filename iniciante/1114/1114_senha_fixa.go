package main

import (
	"fmt"
)

func main() {
	var senha int

	for {
		fmt.Scan(&senha)
		if senha == 2002 {
			fmt.Println("Acesso Permitido")
			break
		}
		fmt.Println("Senha Invalida")
	}
}
