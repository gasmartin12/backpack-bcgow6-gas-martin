/*
Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus
empleados basándose en la cantidad de horas trabajadas por mes
y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20
de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50
de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad
de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/

package main

import (
	"errors"
	"fmt"
)

const (
	salaryCaseC = 1000
	salaryCaseB = 1500
	salaryCaseA = 3000
)

func calcHour(mins float64) (hours float64) {
	return (mins * 1) / 60
}

func calcPercent(category string) (salary float64) {
	switch category {
	case "A":
		return salaryCaseA + (salaryCaseA*50)/100
	case "B":
		return salaryCaseB + (salaryCaseB*50)/100
	case "C":
		return salaryCaseC
	}
	return
}

func addPercent(hours float64, category string) float64 {
	return calcPercent(category)
}

func salary(mins float64, category string) (float64, error) {

	hours := calcHour(mins)
	switch category {
	case "A":
		return addPercent(hours, category), nil
	case "B":
		return addPercent(hours, category), nil
	case "C":
		return addPercent(hours, category), nil
	}
	return 0, errors.New("Error")
}

func main() {
	var mins float64
	var category string

	fmt.Print("Enter the minutes worked in the month: ")
	fmt.Scanf("%g", &mins)
	fmt.Print("Enter your category: ")
	fmt.Scanf("%s", &category)

	salary, err := salary(mins, category)
	if err != nil {
		fmt.Println("Error!!")
	} else {
		fmt.Println("The salary is: ", salary)
	}
}
