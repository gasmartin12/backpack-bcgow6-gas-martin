package main

import "fmt"

/*
un puntero es un tipo de dato, cuyo valor es la direccion de
memoria a la que apunta otro tipo de dato
*/

func increment(pointer *int) {
	*pointer = 30
}

type Thing struct {
	Value int
}

func (pointer *Thing) Update(new int) {
	pointer.Value = new
}

func main() {

	//var pointer *int
	//pointerTwo := new(string)
	var num int = 10
	pointerThree := &num

	fmt.Printf("The direction is: %p\n", &num)
	fmt.Printf("The value is: %d\n", num)
	fmt.Printf("The value is: %d\n", *pointerThree)

	*pointerThree = 20
	fmt.Printf("The value is: %d\n", num)

	increment(pointerThree)
	fmt.Printf("The value is: %d\n", num)

	var t = Thing{
		Value: 50,
	}
	fmt.Printf("The value after is: %d\n", num)

	t.Update(40)
	fmt.Printf("The value before is: %d\n", num)

}
