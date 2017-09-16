package main

import
(
	"fmt"
)

//Le classi in GO vengono rappresentate con le Struct

//Classe Creature

//La classe Creature contiene solo proprietà, 
//quindi solo lo stato dell'oggetto Creature
type Creature struct{
	Name string
	Real bool
}

//Definizione del metodo Dump della classe Creature

func (c Creature) Dump (){
	fmt.Printf("Nome: \t%s\nEsiste:\t%t", c.Name, c.Real)
}

//Definizione della sottoclasse FlyingCreature come figlia di Creature

type FlyingCreature struct{
	Creature
	WingSpan int
}

//Definizione delle Interfacce OOP in GO
type Fooer interface{
	Foo1()
	Foo2()
	Foo3()
}

//Classe astratta
//Gli accessi in GO funzionano in modo poco convenzionale
//Se il nome della classe è in lowercase si tratta 
//di una classe privata; altrimente se è in uppercase
//si tratta di una pubblica

type foo struct{

}

func (f foo) Foo1(){
	fmt.Println("Chiamata Foo1")
}
func (f foo) Foo2(){
	fmt.Println("Chiamata Foo2")
}
func (f foo) Foo3(){
	fmt.Println("Chiamata Foo3")
}

func NewFoo() Fooer{
	return &foo{}
}

func main() {
	dragon := &FlyingCreature {
		Creature{
			"Dragon", false } ,
			15,
	}

	dragon.Dump()

}