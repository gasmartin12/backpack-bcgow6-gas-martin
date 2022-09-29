package main

import "fmt"

const (
	AdditionOperator       = "+"
	SubtractionOperator    = "-"
	MultiplicationOperator = "*"
	DivisionOperator       = "/"
)

func addition(a, b float64) float64 {
	return a + b
}

func substraction(a, b float64) float64 {
	return a - b
}

func multiplication(a, b float64) float64 {
	return a * b
}

func division(a, b float64) (resultado float64) {
	if b == 0 {
		return 0
	}

	return a / b
}

func getOperationFunction(operator string) func(float64, float64) float64 {
	switch operator {
	case AdditionOperator:
		return addition
	case SubtractionOperator:
		return substraction
	case MultiplicationOperator:
		return multiplication
	case DivisionOperator:
		return division
	}

	return nil
}

func main() {
	var returnedOperationFunction func(float64, float64) float64

	returnedOperationFunction = getOperationFunction(SubtractionOperator)

	fmt.Println(returnedOperationFunction(10, 2))
}
