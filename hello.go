package main

import ("fmt")

type Gorra struct{
	marca string
	color string
	precio float32
	plana bool
}

func main(){

	var gorra_negra =Gorra{
		marca: "nike",
		color: "negro",
		precio: 24.5,
		plana: false}
	fmt.Println(gorra_negra)	

	var number1 float32 =14
	var number2 float32 =6

	fmt.Print("la suma es ");
	fmt.Println(number1 + number2);

	fmt.Print("la resta es ");
	fmt.Println(number1 - number2);

	fmt.Print("la producto es ");
	fmt.Println(number1 * number2);

	fmt.Print("la division es ")
	fmt.Println(number1 / number2)

	/*fmt.Print("el resto es ")
	fmt.Println(number1 % number2) solo con valores enteros*/
}