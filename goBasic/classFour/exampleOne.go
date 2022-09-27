package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Perro struct {
	Nombre string   `json:"name"`
	Edad   int      `json:"edad"`
	Raza   string   `json:"raza"`
	Peso   float64  `json:"peso"`
	Color  []string `json:"color"`
	Owner  Dueño    `json:"owner"`
}

func (p Perro) Ladrar() {
	fmt.Println("Guau!")
}

func (p Perro) Dormir() {
	fmt.Println("Estoy durmiendo, soy un perro")
}

func (p Perro) Saludar() {
	fmt.Println("Guau! Soy", p.Nombre)
}

func (p *Perro) Renombrar(new string) {
	p.Nombre = new
	fmt.Println("Guau! Ahora soy", p.Nombre)
}

func (p Perro) Comer() {
	fmt.Println("Estoy comiendo, soy un perro")
}

type Dueño struct {
	Documento        int
	Nombre           string
	NumeroDeContacto int
}

type Gato struct {
	Nombre   string  `json:"name"`
	Edad     int     `json:"edad"`
	Raza     string  `json:"raza"`
	Peso     float64 `json:"peso"`
	Caracter string  `json:"caracter"`
}

func (g Gato) Dormir() {
	fmt.Println("Estoy durmiendo, soy un gato")
}

func (g Gato) Comer() {
	fmt.Println("Estoy comiendo, soy un gato")
}

type Animal interface {
	Dormir()
	Comer()
}

func NewAnimal(tipo string) Animal {
	if tipo == "Perro" {
		return &Perro{}
	} else if tipo == "Gato" {
		return &Gato{}
	}
	return nil
}

type ListaHeterogenea struct {
	Datos []interface{}
}

func main() {
	d1 := Dueño{
		Documento:        1234,
		Nombre:           "Nicolas",
		NumeroDeContacto: 532563234,
	}

	dog := Perro{
		Nombre: "Firulais",
		Edad:   2,
		Raza:   "Pastor",
		Peso:   10.0,
		Color:  []string{"Negro", "Blanco"},
		Owner:  d1,
	}

	// fmt.Printf("Dog: %+v", dog)

	// dog.Owner.Documento = 1234

	// fmt.Printf("\nOwner.Documento: %v", dog.Owner.Documento)

	_, err := json.Marshal(&dog)
	if err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Printf("\njson.Marshal: %+v", &dog)

	dogR := reflect.TypeOf(dog)

	for i := 0; i < dogR.NumField(); i++ {
		// field := dogR.Field(i)
		// fmt.Printf("\nCampo: %s, tiene %s\n", field.Name, field.Tag.Get("json"))
	}

	// dog.Ladrar()
	// dog.Saludar()
	fmt.Printf("\nNombre Antes %s\n", dog.Nombre)
	dog.Renombrar("Cacho")
	fmt.Printf("\nNombre Despues %s\n", dog.Nombre)

	a := NewAnimal("Gato")
	a.Comer()
	a.Dormir()

	lista := ListaHeterogenea{}

	lista.Datos = append(lista.Datos, 1)
	lista.Datos = append(lista.Datos, 5.6)
	lista.Datos = append(lista.Datos, "Hola")

	fmt.Println(lista.Datos)

	e, ok := lista.Datos[0].(int)
	fmt.Println(ok)
	fmt.Println(e)
}
