package ordenar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenarAsc(t *testing.T) {
	s := []int{4, 2, 3, 1}
	esperado := []int{1, 2, 3, 4}
	resultado := Ordenar(s)
	assert.Equal(t, esperado, resultado, "debe ser iguales")
}
