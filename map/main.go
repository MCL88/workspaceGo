package main

import(
	"fmt"
)

//Le mappe sono i "dizionari" di GO, cioè una tabella
//con elementi composti da "chiavi/valori"

//X è una mappa dove la chiave è di tipo stringa e il 
//valore è di tipo intero



func main() {
	//func1()
	tavola_periodica()
}

func func1() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
}

func tavola_periodica(){
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"


	fmt.Println(elements["Li"])
//Un altro modo per descrivere una condizione?
	if name, ok := elements["Un"]; ok{
		fmt.Println(name)
	}
}

//Sto creando una mappa di stringhe che si associano ad altre
//mappe di stringhe che si associano altre stringhe
func tavola_periodica_2() {
	elements := map[string]map[string]string{
		"H" : map[string]string = {
			"name" : "Hydrogen",
			"state" : "gas"
		},
		"He" : map[string]string = {
			"name" : "Helium",
			"state" : "gas"
		},
		"Li" : map[string]string = {
			"name" : "Lithium",
			"state" : "solid"
		},
	}	
}