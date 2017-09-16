package main

import(
	"fmt"
)

//semplice definizione di slice in GO

//Gli slice in GO, adifferenza degli array, permettono
//la modifica sulla loro lunghezza
var x []float64



func main() {
	fmt.Println(x)
	// func1()
	// func2()
	// slice_append()
	// slice_copy()
	// esercizio2();
	min_2()
}

func func1() {
//creo uno slice iniziale di lunghezza 5
	y := make([]float64, 5,10)
	fmt.Println(y)
}

func func2() {
	arr := []float64{1,2,3,4,5}
	r := arr[1:3]
	fmt.Println(r) 
}

func slice_append(){
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4,5)
	fmt.Println(slice1, slice2)
}

//funzionamento slice_copy
func slice_copy(){
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func esercizio1(){
	x := make([]int, 3, 9)
	fmt.Println(len(x))
}

func esercizio2(){
	x := [6]string{"a", "b", "c", "d", "e","f"}
	fmt.Println(x[2:5])
}

func min(){
	x := []float32{
		34, 12, 34, 34,
		21, 43, 1, -340132,
		-1, -1e9,
	}

	var min float32 = 0

	for i := 0; i < len(x); i++{
		if min > x[i] {
			min = x[i]
		}
	}

	fmt.Println(min)

}

//seconda versione
func min_2(){
	x := []float32{
		34, 12, 34, 34,
		21, 43, 1, -340132,
		-1, -1e9,
	}

	var min float32 = 0

	for _, value := range(x){
		if min > value {
			min = value
		}
	}

	fmt.Println(min)

}

