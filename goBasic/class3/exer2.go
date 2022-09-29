/*
Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por alumno)
de sus calificaciones. Se solicita generar una función en
la cual se le pueda pasar N cantidad de enteros y devuelva el
promedio y un error en caso que uno de los números ingresados sea negativo
*/

package main

import (
	"errors"
	"fmt"
)

func calcAverage(ratings ...float64) (float64, error) {
	fmt.Println(ratings)
	var average float64
	for _, qualification := range ratings {
		if qualification > 0 {
			average += float64(qualification)
		} else {
			return 0, errors.New("Invalid qualification")
		}
	}
	return average / float64(len(ratings)), nil
}

func main() {

	average, err := calcAverage(1, 2, 3, -4, 5, 6, 7, 8, 9, 10)
	if average == 0 && err != nil {
		fmt.Println(errors.New("Invalid qualification"))
	} else {
		fmt.Println("The average is:", average)
	}
}
