/*
Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto,
con la información de productos comprados, separados por punto y
coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	message := "Id: 2\t Product: a\t Price: 1\t Count: 1;\n"
	file, err := os.OpenFile("Products.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	if _, err := file.Write([]byte(message)); err != nil {
		panic(err)
	}

	if err := file.Close(); err != nil {
		panic(err)
	}
	fmt.Println("File created succesfully")
}
