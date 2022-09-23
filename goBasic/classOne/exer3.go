package main

import "fmt"

func main() {
	// originals
	// var 1nombre string
	// var apellido string
	// var int edad
	// 1apellido := 6
	// var licencia_de_conducir = true
	// var estatura de la persona int
	// cantidadDeHijos := 2

	// correct
	var nombre string               // do not use numbers in names
	var apellido string             // good
	var edad int                    // first name then type
	apellido_mal := 6               // already declared, do not use numbers in names, more descriptive name
	var licencia_de_conducir = true //good
	var estatura_de_la_persona int  //use _ to separate words
	cantidadDeHijos := 2            // good

	fmt.Println(nombre, apellido, edad, apellido_mal, licencia_de_conducir, estatura_de_la_persona, cantidadDeHijos)
}
