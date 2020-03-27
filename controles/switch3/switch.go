package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "FLOAT"
	case string:
		return "string"
	case func():
		return "Função"
	default:
		return "não sei"
	}
}

func main() {
	fmt.Println(tipo(1))
	fmt.Println(tipo(2.54))
	fmt.Println(tipo("Marcelo"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(true))
}
