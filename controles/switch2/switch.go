package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("BOM DIA")
	case t.Hour() < 18:
		fmt.Println("BOA TARDE")
	default:
		fmt.Println("BOA NOITE")
	}

}
