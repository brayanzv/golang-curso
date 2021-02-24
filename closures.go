package main

import(
	"fmt"
)

func main(){
fmt.Println(gorras(45, "$"))

fmt.Println(gorras(24, "$"))
}

func gorras(pedido float32, moneda string) (string, float32, string){

	precio:= func() float32{
		return pedido *7
	}

	return "el numero de gorras pedidas es ", precio(), moneda
}