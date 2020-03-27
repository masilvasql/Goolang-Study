package main

import "fmt"

func imprimirResultado(nota float64) {
	// GO não aceita parenteses no IF
	if nota >= 7 {
		fmt.Println("Aprovado com nota: ", nota)
	} else {
		fmt.Println("Reprovado com a nota ", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(6)
}
