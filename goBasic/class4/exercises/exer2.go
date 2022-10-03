package main

import "fmt"

type Matrix struct {
	Values    []float64
	Height    int
	Width     int
	quadratic bool
	MaxValue  float64
}

func main() {
	matrix := NewMatrix(1, 2, 3, 4, 5, 6, 7, 8)
	matrix.Print()
}

func NewMatrix(height int, width int, values ...float64) (matrix Matrix) {
	matrix.Height = height
	matrix.Width = width
	matrix.Values = values

	matrix.quadratic = false
	if height == width {
		matrix.quadratic = true
	}

	actualMax := matrix.Values[0]
	for _, value := range values {
		if value > actualMax {
			actualMax = value
		}
	}
	matrix.MaxValue = actualMax
	return
}

func (matrix Matrix) Print() {
	num := 0

	for i := 0; i < matrix.Height; i++ {
		for j := 0; j < matrix.Width; j++ {
			fmt.Printf("%f ", matrix.Values[num+j])
		}
		num = (i + 1) * matrix.Width
		fmt.Print("\n")
	}
}
