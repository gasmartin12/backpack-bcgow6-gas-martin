/*
Ejercicio 3 - Impuestos de salario #3

Repite el proceso anterior, pero ahora implementando
"fmt.Errorf()", para que el mensaje de error reciba por
parámetro el valor de "salary" indicando que no alcanza
el mínimo imponible (el mensaje mostrado por consola deberá
decir: "error: el mínimo imponible es de 150.000 y el
salario ingresado es de: [salary]", siendo [salary] el
valor de tipo int pasado por parámetro)
*/

package main

import (
	"fmt"
	"os"
)

const salaryMin = 150000

var (
	msg      = "You must pay tax"
	errorMsg = "error: the minimum taxable amount is %d and the salary entered is: [%d]"
)

func main() {
	salary := 20000

	if salary < salaryMin {
		err := fmt.Errorf(errorMsg, salaryMin, salary)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(msg)
}
