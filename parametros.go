package main

import("fmt")

func main(){
pantalon("rojo", true)
pantalons("rojo", "largo", "sin bolsillos", "nike")
}

func pantalons(caracteristicas ...string){
	for _,caracteristica := range caracteristicas{
		fmt.Println(caracteristica)
	}
}

func pantalon(color string, es_corto bool){
	fmt.Println(color)
	fmt.Println(es_corto)
}