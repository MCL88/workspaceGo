package main

import
(
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perim() float64
}

type Rect struct{
	base, altezza float64
}

type Circle struct{
	raggio float64
}

func(r Rect) area() float64 {
	return r.base * r.altezza
}

func(r Rect) perim() float64 {
	return (r.base + r.altezza)*2
}

func(c Circle) area() float64 {
	return math.PI * c.raggio * c.raggio
}

func(c Circle) perim() float64 {
	return 2 * math.PI * c.raggio
}

func misura(g geometry)
{
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() 
{
	c := Circle{raggio:5}
	misura(c)
}