/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de
sus empleados al momento de depositar el sueldo, para cumplir
el objetivo es necesario crear una función que devuelva el
impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se
le descontará un 17% del sueldo y si gana más de $150.000 se
le descontará además un 10%.
*/
package main

import "fmt"

const (
	firstSalary      = 50000
	secondSalary     = 150000
	firstPercentage  = 17
	secondPercentage = 27
)

func calcPercent(salary, percen int) int {
	return (salary * percen) / 100

}

func main() {

	tax := calcPercent(firstSalary, firstPercentage)
	taxTow := calcPercent(secondSalary, secondPercentage)

	fmt.Println("The tax for employees earning more than ", firstSalary, " is:", tax)
	fmt.Println("The tax for employees earning more than ", secondSalary, " is:", taxTow)
}
