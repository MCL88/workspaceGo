package main 
import "fmt"

var i int = 0

func main() {
		
	for {

		fmt.Println(i)
		i++
		if(i == 100) {
			break
		}
		
	}
}