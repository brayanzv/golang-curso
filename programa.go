package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("*****Mi primer programa en go*****")
	fmt.Println("hola " + os.Args[1] + "bienvenido al programa en go")
	edad, _ /*err*/ := strconv.Atoi(os.Args[2])
	//fmt.Println(err, "hay un error")
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad o mucho mayor")
	}

	numero, _ := strconv.Atoi(os.Args[3])

	if numero%2 == 0 {
		fmt.Println("el numero es par")
	} else {
		fmt.Println("el numero es impar")
	}

	peliculas := []string{"pelicula2", "peliculas1", "peliculas3"}

	for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}

	//momento := time.Now()

	hoy := time.Now().Weekday()
	fmt.Println(hoy)
	fmt.Println(time.Now())
	switch hoy {
	case 0:
		fmt.Println("hoy es domingo")
	case 1:
		fmt.Println("hoy es lunes")
	case 2:
		fmt.Println("hoy es martes")
	case 3:
		fmt.Println("hoy es miercoles")

	case 4:
		fmt.Println("hoy es jueves")

	case 5:
		fmt.Println("hoy es viernes")

	}
}
