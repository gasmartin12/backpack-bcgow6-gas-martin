package operaciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// los test se ejecutan desde la carpeta donde esta contenido
// simpre va con la palabra test adelante
// go test -v para ejecutar los test

func TestAddGood(t *testing.T) {
	//Arrange -> que son los valores de entrada y que espero que devuelva
	num1 := 10
	num2 := 5
	esperado := 15
	//Act -> llamado a la funcion
	resultado, _ := Add(num1, num2)
	//Asset
	assert.Equal(t, esperado, resultado, "deben ser iguales")
}

func TestAddBad(t *testing.T) {
	//Arrange -> que son los valores de entrada y que espero que devuelva
	num1 := 1

	num2 := 5
	//Act -> llamado a la funcion
	_, err := Add(num1, num2)
	//Asset
	assert.Nil(t, err)
}

// exer 1
func TestSub(t *testing.T) {
	//Arrange
	num1 := 10
	num2 := 5
	esperado := 5
	//Act -> llamado a la funcion
	resultado := Sub(num1, num2)
	//Asset
	assert.Equal(t, esperado, resultado, "deben ser iguales")
}

// exer 3
func TestDivide(t *testing.T) {
	var num1 float64 = 4
	var num2 float64 = 2
	var esperado float64 = 2
	resultado, _ := Divide(num1, num2)
	assert.Equal(t, esperado, resultado, "deben ser iguales")
}

func TestDivideError(t *testing.T) {
	var num1 float64 = 4
	var num2 float64 = 0
	_, err := Divide(num1, num2)
	assert.Nil(t, err, "Error encontrado: %s", err)
}
