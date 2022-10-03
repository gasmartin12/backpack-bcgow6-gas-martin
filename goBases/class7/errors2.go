package main

import (
	"errors"
	"fmt"
	"os"
)

func division(num1, num2 float64) (result float64, err error) {
	if num2 == 0 {
		err = errors.New("no se puede dividir por 0")
		return
	}
	result = num1 / num2
	return
}
func calcular(op string, num1, num2 float64) (result float64, err error) {
	switch op {
	case "/":
		result, err = division(num1, num2)
	default:
		err = errors.New("operador no soportado")
	}
	return
}
func main() {
	result, err := calcular("+", 10, 2)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("resultado: %f", result)
}
