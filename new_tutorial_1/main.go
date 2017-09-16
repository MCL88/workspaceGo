package main

import "fmt"

var scelta int = -1

const (
	a = iota
	b
)

func unhex(c byte) byte {
	fmt.Printf("\n %d", c - 1)
	switch{
	case '0' <= c && c <= '9' :
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func main() {
	fmt.Printf("%d %d\n", a, b)
	for scelta != 0 {
		fmt.Printf("Inserire un carattere qualsiasi(Solo invio per chiudere il programma): ")
		fmt.Scanf("&c", &scelta)
		if(scelta != 0) {
			fmt.Printf("\n %c\n",unhex(byte(scelta)))
			scelta = -1
		}else {
			fmt.Printf("Ciao!")
		}

	}
}

