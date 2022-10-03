/*
Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado,
para ello requiere que: se imprima por pantalla mostrando los
valores tabulados, con un título (tabulado a la izquierda para el ID
y a la derecha para el Precio y Cantidad), el precio, la cantidad
y abajo del precio se debe visualizar el total
(Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./data.csv")
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		line := fscanner.Text()
		lineSlice := strings.Split(line, ";")
		id := lineSlice[0]
		price := lineSlice[1]
		quantity := lineSlice[2]
		fmt.Printf("%s\t %s\t %s\n", id, price, quantity)
	}
}
