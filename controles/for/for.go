package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // não tem while em go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par ", i)
		} else {
			fmt.Println("Impar ", i)
		}
	}

	for {
		//Laço infinito
		fmt.Println("Infinito")
		time.Sleep(time.Second)
	}

}
