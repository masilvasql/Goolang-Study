package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //quem conta é o compilador o array

	//retorna o indice e o elemento
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, numero := range numeros {
		fmt.Printf("%d\n", numero)
	}
}
