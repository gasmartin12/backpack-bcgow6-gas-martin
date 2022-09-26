package main

import "fmt"

const (
	AdditionOperator       = "+"
	SubstractionOperator   = "-"
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

func division(a, b float64) float64 {
	if b == 0 {
		return 0
	}

	return a / b
}

func orchestrateOperation(numbers []float64, operation func(num1, num2 float64) float64) float64 {
	var result float64

	for index := range numbers {
		if index == 0 {
			result = numbers[index]
		} else {
			result = operation(result, numbers[index])
		}
	}

	return result
}

func calculate(operator string, values ...float64) float64 {
	switch operator {
	case AdditionOperator:
		return orchestrateOperation(values, addition)
	case SubstractionOperator:
		return orchestrateOperation(values, substraction)
	case MultiplicationOperator:
		return orchestrateOperation(values, multiplication)
	case DivisionOperator:
		return orchestrateOperation(values, division)
	}
	return 0
}

func main() {
	fmt.Println(
		calculate(AdditionOperator, 10, 20, 40),
		calculate(SubstractionOperator, 5, 5),
		calculate(MultiplicationOperator, 2, 4, 8),
		calculate(DivisionOperator, 10, 20, 30),
	)

	fmt.Println(calculate(SubstractionOperator, 5, 5))
	fmt.Println(calculate(MultiplicationOperator, 2, 4, 8))
	fmt.Println(calculate(DivisionOperator, 10, 20))
	fmt.Println(calculate(DivisionOperator, 10, 20, 30))
}
