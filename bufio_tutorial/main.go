package main

import
(
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple shell")
	fmt.Println("------------------------")

	for{
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		//converte il \n in carattere vuoto
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0{
			fmt.Println("Ma sei scemo?")
		}
	}
}