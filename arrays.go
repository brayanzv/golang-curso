package main

import "fmt"

func main(){
	var peliculas [3] string;
	peliculas[0]="la verdad duele"
	peliculas[1]="ciudadano ejemplar"
	peliculas[2]="el gran torino"

	//fmt.Println(peliculas)
	fmt.Println(peliculas[1])
	array2()
	
}
//esta es otra menera de definir arrays
func array2(){
	var peliculas = [3]string {
	"la verdad duele",
	"ciudadano ejemplar",
	"gran torino"}

	fmt.Println(peliculas[2])
}