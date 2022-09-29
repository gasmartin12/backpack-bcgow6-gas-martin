/*
Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan
calcular algunas estadísticas de calificaciones de los alumnos
de un curso, requiriendo calcular los valores mínimo, máximo
y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo
se quiere realizar (mínimo, máximo o promedio) y que devuelva otra
función ( y un mensaje en caso que el cálculo no esté definido) que
se le puede pasar una cantidad N de enteros y devuelva el cálculo
que se indicó en la función anterior

Ejemplo:
const (

	minimum = "minimum"
	average = "average"
	maximum = "maximum"

)

...

minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...

minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/
package main

import (
	"errors"
	"fmt"
)

func minFunc(ratings ...float64) (min float64) {
	if len(ratings) > 0 {
		min = ratings[0]
	}
	for _, qualification := range ratings {
		if qualification < min {
			min = qualification
		}
	}
	return
}

func averageFunc(ratings ...float64) (average float64) {

	for _, qualification := range ratings {
		if qualification > 0 {
			average += float64(qualification)
		}
	}
	average = average / float64(len(ratings))
	return
}

func maxFunc(ratings ...float64) (max float64) {
	if len(ratings) > 0 {
		max = ratings[0]
	}
	for _, qualification := range ratings {
		if qualification > max {
			max = qualification
		}
	}
	return
}

func whatCalc(calc string) (float64, error) {

	switch calc {
	case "min":
		return minFunc(2, 3, 3, 4, 10, 2, 4, 5), nil
	case "ave":
		return averageFunc(2, 3, 3, 4, 1, 2, 4, 5), nil
	case "max":
		return maxFunc(2, 3, 3, 4, 1, 2, 4, 5), nil
	}
	return 0, errors.New("invalid calc")
}

func main() {

	var calculation string
	fmt.Println("Select the calculation")
	fmt.Println("Min Value: m")
	fmt.Println("Average Value: a")
	fmt.Println("Max Value: x ")

	fmt.Scanf("%s", &calculation)

	switch calculation {
	case "m":
		result, err := whatCalc("min")
		if err != nil {
			fmt.Println(errors.New("invalid calc"))
		} else {
			fmt.Println("The result for min calc is: ", result)
		}
	case "a":
		result, err := whatCalc("ave")
		if err != nil {
			fmt.Println(errors.New("invalid calc"))
		} else {
			fmt.Println("The result for average calc is: ", result)
		}
	case "x":
		result, err := whatCalc("max")
		if err != nil {
			fmt.Println(errors.New("invalid calc"))
		} else {
			fmt.Println("The result for max calc is: ", result)
		}
	default:
		fmt.Println(errors.New("Invalid option"))
	}
}
