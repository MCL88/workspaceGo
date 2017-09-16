package main
import(
	"os"
	"fmt"
	"strings"
	"testing"
)

var s, sep string

func main() {
	
	fmt.Println("Stai eseguendo il programma", os.Args[0])
	//echo1()
	//echo2()
	//echo3()
	echo2_1()
	fmt.Println(s)
}

func echo1() {

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

}

func echo2() {

//L'underscore rappresenta il blank value in GO. 
//range è una funzione che restituisce l'indice e l'oggetto di un array
//il blank value non fa altro che prendere l'indice ed eliminarlo dalla memoria perché inutile
	for _, arg := range(os.Args[1:]){
		s += sep + arg
		sep = " "
	}

}

func echo2_1() {

//L'underscore rappresenta il blank value in GO. 
//range è una funzione che restituisce l'indice e l'oggetto di un array
//il blank value non fa altro che prendere l'indice ed eliminarlo dalla memoria perché inutile
	for i, arg := range(os.Args[1:]){
		fmt.Println(i, arg)
	}

}

func echo3(){

	fmt.Println(strings.Join(os.Args[1:], " "))
}