package main

import "fmt"

const (
	PEQ = "Pequeño"
	MED = "Mediano"
	GRA = "Grande"
)

// Producto
type Producto interface {
	CalcularCosto() float64
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

func main() {
	zapatilla := nuevoProducto(PEQ, "zapatilla", 2_000)
	mesa := nuevoProducto(GRA, "mesa", 15_000)
	tienda := tienda{}
	tienda.Agregar(zapatilla)
	tienda.Agregar(mesa)
	total := tienda.Total()
	fmt.Println(total)
}

func nuevoProducto(tipo string, nombre string, precio float64) (p producto) {
	p.tipo = tipo
	p.nombre = nombre
	p.precio = precio
	return
}

func (p producto) CalcularCosto() (costoTotal float64) {
	switch p.tipo {
	case PEQ:
		costoTotal = p.precio
	case MED:
		costoTotal = p.precio * 1.03
	case GRA:
		costoTotal = (p.precio * 1.06) + 2500
	}
	return
}

// Tienda
type Ecommerce interface {
	Total() float64
	Agregar(p producto)
}

type tienda struct {
	productos []producto
}

func nuevaTienda() (e Ecommerce) {
	return
}

func (t tienda) Total() (precioTotal float64) {
	for _, producto := range t.productos {
		precioTotal += producto.CalcularCosto()
	}
	return
}

func (t *tienda) Agregar(p producto) {
	t.productos = append(t.productos, p)
	return
}
