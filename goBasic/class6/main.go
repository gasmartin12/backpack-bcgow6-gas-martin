package main

import "fmt"

func main() {

	//var puntero *int
	//p2 := new(int)

	var numero int = 10

	p3 := &numero

	fmt.Printf("El valor de numero es %d\n", numero)

	fmt.Printf("El valor del puntero es %p\n", p3)

	//Incrementar(p3)

	//fmt.Printf("El valor de numero es %d\n", numero)

	//var c = Coso{Valor: 20}

	//fmt.Printf("Coso antes de actualizar:%+v", c)

	//c.Actualizar(10)

	//fmt.Printf("Coso despues: %+v", c)

}

func Incrementar(puntero *int) {
	fmt.Printf("Antes %d\n", puntero)
	*puntero++
	fmt.Printf("Despues %d\n", puntero)
}

type Coso struct {
	Valor int
}

func (c *Coso) Actualizar(new int) {
	c.Valor = new
}
