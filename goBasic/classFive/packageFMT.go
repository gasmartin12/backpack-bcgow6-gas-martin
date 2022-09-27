package main

import (
	"fmt"
)

func main() {
	messageHello := "Hello world :) \n"
	messageBye := "Bye world :( \n"
	b, err := fmt.Print(messageHello, messageBye)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(b)

	flag := true
	fmt.Printf("la variable es %b", flag)

	num := 1222.123123
	fmt.Printf("El valor total es: %f\n", num)
	fmt.Printf("EL valor cortado es: %10.2f\n", num)

	name, edad := "Gas", 23

	fmt.Println("Hola, ", name, "! Tengo", edad, " años")
	// ==
	fmt.Printf("Hola mi nombre es %s y tengo %d años", name, edad)
}
