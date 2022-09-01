package main

import "fmt"

type Item struct {
	ProdutoID int
	Qtde      int
	Preco     float64
}

type Pedido struct {
	UserID int
	Itens  []Item
}

func (p Pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.Itens {
		total += item.Preco * float64(item.Qtde)
	}

	return total
}

func main() {

	pedido := Pedido{
		UserID: 1,
		Itens: []Item{
			{
				ProdutoID: 1,
				Qtde:      2,
				Preco:     10,
			},
			{
				ProdutoID: 2,
				Qtde:      5,
				Preco:     10,
			},
		},
	}

	fmt.Printf("O Valor total do pedido é R$ %.2f", pedido.valorTotal())

}
