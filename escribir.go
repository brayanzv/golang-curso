package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("***Escribir en el Fichero")

	escribir := os.Args[1] + "\n"
	//escribir :=ioutil.WriteFile("texto.txt", escribir, 0777)// agrega escritura pero esta elimina la anterior en el archivo
	fichero, err := os.OpenFile("texto.txt", os.O_APPEND, 0777)
	showError(err)

	FileWrite, err := fichero.WriteString(escribir)
	fmt.Println(FileWrite)
	showError(err)

	texto, errorDeFichero := ioutil.ReadFile("texto.txt")
	showError(errorDeFichero)

	fmt.Println(string(texto))

}
func showError(e error) {
	if e != nil {
		panic(e)
	}
}
