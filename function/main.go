package main 

import
(
    "fmt"
)

func hi(name string) {
	fmt.Println("Ciao", name + "!")
}

func main() {
	hi("Michele")
}