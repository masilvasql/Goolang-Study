package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiro
	// Um ponteiro é uma referência de memória
	// * representa o ponteiro

	var p *int = nil //criado um ponteiro

	p = &i //pega o endereço de memória da variável

	*p++ // (desreferenciando) pega o valor associado ao ponteiro
	i++

	fmt.Println(p, i, *p, &i) //p = endereço de memória de i, i = valor de i, *p = valor de i, &i, endereço de memória de i

}
