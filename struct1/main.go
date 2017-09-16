package main

import(
	"fmt" 
	"math"
)

func distance(x1, y1, x2, y2 float64) float64{
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}


type Circle struct{
	x,y,r float32
}

func (c *Circle) area() float32{
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64{
	 l := distance(r.x1, r.y1, r.x1, r.y2)
	 w := distance(r.x1, r.y1, r.x2, r.y1)
	 return l * w
} 



func main() {
	r := new(Rectangle)
	c := new(Circle)

	c = &Circle{3,5,7}
	r = &Rectangle{0,0,10,10}

	fmt.Println(c.x)
	fmt.Println(c.y)
	fmt.Println(c.r)
	fmt.Println(c.area())
	fmt.Println(r.area())
}