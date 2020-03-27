package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("SOMA =>", a+b)
	fmt.Println("SUBTRAÇÃO =>", a-b)
	fmt.Println("DIVISÃO =>", a/b)
	fmt.Println("MULTIPLICAÇÃO =>", a*b)
	fmt.Println("MÓDULO =>", a%b)

	//BITWISE OPERADORES BINÁRIOS
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("OU =>", a|b)  // 11 | 10 = 11
	fmt.Println("XOR =>", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	//outras operações usando MATH

	fmt.Println("MAIOR => ", math.Max(float64(a), float64(b)))
	fmt.Println("MENOR => ", math.Min(c, d))
	fmt.Println("MENOR => ", math.Pow(c, d)) //potência 3 ^ 2

}
