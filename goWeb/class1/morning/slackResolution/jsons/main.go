package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre string
	Edad   int64
}

func main() {
	p := &Persona{
		Nombre: "Juan",
		Edad:   22,
	}

	// Marshal -> ayudarnos a transformar nuestra data a un formato json
	pBytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	// Devolver como un json
	fmt.Println(string(pBytes))

	// Unmarshal
	p1 := &Persona{}
	err = json.Unmarshal(pBytes, p1)
	if err != nil {
		panic(err)
	}

	fmt.Println(p1) // stuct Persona -> con los datos pasados desde la función unmarshal
	fmt.Println("Nombre ", p1.Nombre)
	fmt.Println("Edad ", p1.Edad)

	fmt.Printf("%#v", p1)
	fmt.Printf("%+v", p1)
}
