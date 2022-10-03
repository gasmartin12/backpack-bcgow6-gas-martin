package main

import "fmt"

func main() {
	// originals
	// var apellido string = "Gomez"
	// var edad int = "35"
	// boolean := "false";
	// var sueldo string = 45857.90
	// var nombre string = "Julián"

	// correct
	var apellido string = "Gomez"  // good
	var edad int = 35              // is an integer, without "
	boolean := false               // is an bool, without "
	var sueldo string = "45857.90" //is an string, with "
	var nombre string = "Julián"   // good

	fmt.Println(apellido, edad, boolean, sueldo, nombre)
}
