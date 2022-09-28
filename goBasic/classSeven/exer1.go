/*
Ejercicio 1 - Impuestos de salario #1
En tu función "main", define una variable llamada "salary"
y asignarle un valor de tipo "int".
Crea un error personalizado con un struct que implemente
"Error()" con el mensaje "error: el salario ingresado no
alcanza el mínimo imponible" y lánzalo en caso de que "salary"
sea menor a 150.000. Caso contrario, imprime por consola el
mensaje "Debe pagar impuesto".
*/

package main

import (
	"fmt"
	"os"
)

const salaryMin = 150000

var (
	msg      = "you must pay tax"
	errorMsg = "the salary paid does not the taxable minimum"
)

type SalaryError struct {
	StatusCode int
	Message    string
}

func (err *SalaryError) Error() string {
	return err.Message
}

func getSalaryError(salary int) (code int, err error) {
	if salary < salaryMin {
		return 500,
			&SalaryError{
				StatusCode: 500,
				Message:    errorMsg,
			}
	}
	return 200, nil
}

func main() {
	salary := 20000
	code, err := getSalaryError(salary)
	if code != 200 {
		fmt.Printf("error: %d \n %s\n", code, err)
		os.Exit(1)
	}
	fmt.Println(msg)

}
