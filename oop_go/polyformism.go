package main

import "fmt"

//Classe Creature

type Creature struct{
	Name string
	Real bool
}

//Due modi per implementare la funzione Dump
//Questo è il concetto di Poliformismo

func Dump(c * Creature){
	fmt.Printf("Nome: \t%s\nEsiste:\t%t", c.Name, c.Real)
	fmt.Println()
	fmt.Println()
}

func (c Creature) Dump(){
	fmt.Printf("Nome: \t%s\nEsiste:\t%t", c.Name, c.Real)
	fmt.Println()
	fmt.Println()
}

//Classe FlyingCreature

type FlyingCreature struct{
	Creature
	WingSpan int
}

func (fc FlyingCreature) Dump(){
	fmt.Printf("Nome: \t\t%s\nEsiste:\t\t%t\nApertura ali:\t",
				fc.Name, 
				fc.Real,
				fc.WingSpan)
}

type Unicorn struct{
	Creature
}

type Dragon struct{
	FlyingCreature
}

type Pterodactyl struct{
	FlyingCreature
}

func NewPterodactyl(wingspan int) *Pterodactyl{
	pet := &Pterodactyl{
		FlyingCreature{
			Creature{
				"Pterodactyl",
				true,
			},
			wingspan,
		},
	}
	return pet
}

type Dumper interface{
	Dump()
}

func main() {
	creature := &Creature{
		"Creatura",
		false,
	}

	creature.Dump()
	fmt.Println()
	Dump(creature)
}