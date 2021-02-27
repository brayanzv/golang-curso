package main

import "fmt"
func main(){
	peliculas:= []string{
		"el gran torino",
		"duro de matar",
		"piratas del caribe",
	}
	peliculas = append(peliculas, "camp rock", "daniel boon")

	fmt.Println(len(peliculas))//me indica cuantos elementos tengo dentrode la array/elements have into the array
	fmt.Println(peliculas[0:2])

}