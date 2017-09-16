package main

import (
	"fmt"
)

var rain int = 2 
var a int = 3
var b int = 12

func main() {
	if(rain == 0) {

		fmt.Println("Non sta piovendo!")
	}
	if(rain == 1) {

		fmt.Println("Prendi l'ombrello")
	}
	if(rain == 2) {

		fmt.Println("Meglio che rimani a casa")
	}

	if(a <= 3) {

		if(b >= 8) {

			fmt.Println("a <= 3 e b >= 8")
		}
	}

}