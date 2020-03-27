package main

import "fmt"

/*se os parâmetros da função
forem iguais, pode ser colocado
o tipo somente no último e os outros
serão atribuídos com o mesmo tipo

a função no GO pode retornar mais de um parâmetro
*/

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 != trab2 //ou exclusivo, no GO não tem
	tomarSorvete := trab1 || trab2

	return comprarTV50, comprarTV32, tomarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)

	fmt.Printf("TV 50: %t TV32: %t sorvete: %t saudável %t", tv50, tv32, sorvete, !sorvete)
}
