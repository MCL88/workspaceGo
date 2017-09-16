package main

import (
	
	"fmt"
)

var voto string = "F"

func main() {
	
	switch {

		case voto == "A":
			fmt.Println("Ottimo!")
		case voto == "B":
			fmt.Println("Buono")
		case voto == "C":
			fmt.Println("Sufficente")
		case voto == "D":
			fmt.Println("Scarso!")
		default:
			fmt.Println("SV")
	}

}