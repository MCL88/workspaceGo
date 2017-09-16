package main

import(
	"fmt"
)

func main() {
	var a,b int = 1, 2
	add := func(x, y int) int{
		return x +y
	}

	fmt.Println(add(a,b))
	newIncGen := increment()

	fmt.Println(newIncGen())
	fmt.Println(newIncGen())
	fmt.Println(newIncGen())

	fmt.Println(half(12))
	fmt.Println(add2(1,4,20,23,56,23,5))
}

func increment() func() uint{
	i := uint(0)
	return func()(r uint){
		r = i
		i++
		return
	}

}

func add2(args ...int) int{
	total := 0
	for _, v := range(args){
		total +=v
	}
	return total
}

func half(x int)(even int, isEven bool) {
	even = x % 2
	if(even == 0){
		isEven = true
	}else{
		isEven = false
	}
	return
}