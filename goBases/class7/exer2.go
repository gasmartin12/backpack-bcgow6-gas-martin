/*
Ejercicio 2 - Impuestos de salario #2

Haz lo mismo que en el ejercicio anterior pero reformulando
el código para que, en reemplazo de "Error()",
se implemente "errors.New()".
*/

package main

import (
	"errors"
	"fmt"
	"os"
)

const salaryMin = 150000

var (
	msg      = "you must pay tax"
	errorMsg = "the salary paid does not the taxable minimum"
)

func getSalaryError(salary int) (code int, err error) {
	if salary < salaryMin {
		return 500, errors.New(errorMsg)
	}
	return 200, nil
}

func main() {
	salary := 20000
	code, err := getSalaryError(salary)
	if code != 200 {
		fmt.Printf("error: %d\n %s\n", code, err)
		os.Exit(1)
	}
	fmt.Println(msg)
}
