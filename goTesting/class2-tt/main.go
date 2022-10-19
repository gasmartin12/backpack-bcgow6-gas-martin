package main

import (
	"fmt"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class2-tt/operaciones"
	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class2-tt/ordenar"
)

func main() {
	a, b := 10, 3
	sum, err := operaciones.Add(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)

	s := []int{4, 2, 3, 1}
	ordenar.Ordenar(s)
}
