/*
Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar
para las mascotas. Por el momento solo tienen tarántulas, hamsters,
perros, y gatos, pero se espera que puedan haber muchos más animales
que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo
texto con el animal especificado y que retorne una función y un mensaje
(en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base
a la cantidad del tipo de animal especificado.
*/

package main

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarántula"
)

func animalDog(dogs int) float64 {
	return 10 * float64(dogs)
}

func animalCat(cats int) float64 {
	return 5 * float64(cats)
}

func animalTarantula(tarantula int) float64 {
	return (150 * float64(tarantula)) / 1000
}

func animalHamster(hamster int) float64 {
	return (250 * float64(hamster)) / 1000
}

func Animal(animal string) (func(int) float64, string) {
	switch animal {
	case dog:
		return animalDog, ""
	case cat:
		return animalCat, ""
	case hamster:
		return animalHamster, ""
	case tarantula:
		return animalTarantula, ""
	default:
		return nil, "Invalid animal"
	}
}

func main() {
	var count float64

	animaldog, msg := Animal(dog)
	if msg != "" {
		fmt.Println(msg)
	} else {
		count += animaldog(2)
	}

	animalcat, msg := Animal(cat)
	if msg != "" {
		fmt.Println(msg)
	} else {
		count += animalcat(5)
	}

	animalhamster, msg := Animal(hamster)
	if msg != "" {
		fmt.Println(msg)
	} else {
		count += animalhamster(4)
	}

	animaltarantula, msg := Animal(tarantula)
	if msg != "" {
		fmt.Println(msg)
	} else {
		count += animaltarantula(2)
	}

	fmt.Println("Food", count, "Kg")
}
