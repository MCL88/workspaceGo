package main

import(
	"fmt"
)

func main() {
	//func1()
	//func2()
	//func3()
	func_range()
}

func func1() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func func2() {
	var x [5]float64
	x[0] = 10
	x[1] = 23
	x[2] = 3
	x[4] = 5
	var total float64

	for i := 0; i < 5; i++{
		total += x[i]
	}

	fmt.Println(total/5)
}

func func3() {
	var x [5]float64
	x[0] = 10
	x[1] = 23
	x[2] = 3
	x[4] = 5
	var total float64

	for i := 0; i < len(x); i++{
		total += x[i]
	}

	fmt.Println(total/float64(len(x)))
}

func func_range() {
	var x [5]float64
	x[0] = 10
	x[1] = 23
	x[2] = 3
	x[4] = 5
	var total float64

//range restituisce la coppia indice/valore dell'array
	for _, value := range(x){
		total += value
	}

	fmt.Println(total/float64(len(x)))
}

