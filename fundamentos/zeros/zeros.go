package main

import "fmt"

func main() {
	var a int     //inicia por padrão com 0
	var b float64 //inicia por padrão com 0
	var c bool    //inicia por padrão com false
	var d string  //inicia por padrão com ""
	var e *int    //inicia por padrão com  nil

	fmt.Printf("%v %v %v %q %v ", a, b, c, d, e)
}
