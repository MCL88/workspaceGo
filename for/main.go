package main 

import
(
    "fmt"
)

func main() {
  for i := 0; i <= 10; i = i + 1 {
    fmt.Println(i)
  }

  for i := 0; i <= 10; i = i + 1 {
    for j := 0; j <= 10; j = j + 1 {
       fmt.Printf("(%2d,%2d)", i,j)  
    }
    fmt.Println()
  }

}