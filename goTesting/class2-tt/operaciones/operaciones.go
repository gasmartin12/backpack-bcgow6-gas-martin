package operaciones

import (
	"fmt"
)

func Add(a, b int) (int, error) {
	if a == 0 {
		return 0, fmt.Errorf("a no puede ser %d", a)
	}
	return a + b, nil
}

// exer 1
func Sub(a, b int) int {
	return a - b
}

// exer 3
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("a no puede ser %f", b)
	}
	return a / b, nil

}
