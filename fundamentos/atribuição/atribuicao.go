package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	// := só pode ser utilizado no momento da criação da variável

	i := 3 //inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 3
	i *= 2 // i = i * 3
	i %= 2 // i = i % 3

	fmt.Println(i)

	x, y := 1, 2

	fmt.Println(x, y)

	x, y = y, x

	fmt.Println(x, y)

}
