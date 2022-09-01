package main

import "fmt"

type Imprimivel interface {
	toString() string
}

type Pessoa struct {
	nome      string
	sobrenome string
}

type Produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p Pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p Produto) toString() string {
	return fmt.Sprintf("%s - R$%2.f", p.nome, p.preco)
}

func teste() {

}

func imprmir(x Imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa Imprimivel = Pessoa{
		nome:      "Marcelo",
		sobrenome: "Abrão",
	}

	fmt.Println(coisa.toString())

	coisa = Produto{
		nome:  "Lápis",
		preco: 1.00,
	}

	fmt.Println(coisa.toString())
}
