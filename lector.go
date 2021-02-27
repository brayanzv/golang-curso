package main

import (
		"fmt"
		"io/ioutil"
		//"os"
)

func main(){
	fmt.Println("Lector de Ficheros")

	fichero, errorDeFichero := ioutil.ReadFile("tesxto.txt")  
	showError(errorDeFichero)
	//fmt.Println(errorDeFichero, " holas")
	fmt.Println(string(fichero))
}

func showError(e error){
	if e != nil {
		panic(e)
	}
}