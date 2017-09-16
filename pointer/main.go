package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	
	a := 42

	var meters float64
	fmt.Println("Inserire i metri:")
	fmt.Scan(&meters)
	yards := meters*metersToYards;
	fmt.Printf("%d \n", 100)
	fmt.Printf("%b \n", 100)
	fmt.Printf("%c \n", 44)
	fmt.Printf("%e \n", 100.0)
	fmt.Println(meters, " metri sono pari a: ", yards, " yards.");
	fmt.Println("A è pari a: ", a);
	fmt.Println("meters si trova all'indirizzo di memoria: ", &meters);
	fmt.Println("yards si trova all'indirizzo di memoria: ", &yards);



}