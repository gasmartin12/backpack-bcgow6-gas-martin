/*
Ejercicio 3 - Calcular Precio
Una empresa nacional se encarga de realizar venta de productos,
servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el
precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte
demanda y para optimizar la velocidad requieren que el cálculo de la
sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el
precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el
precio total (precio * media hora trabajada, si no llega a
	trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve
el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar
por pantalla el monto final (sumando el total de los 3).
*/

package main

import (
	"errors"
	"fmt"
)

type Products struct {
	Name  string
	Price float64
	count int
}
type Services struct {
	Name      string
	Price     float64
	MinWorked float64
}
type Maintenance struct {
	Name  string
	Price float64
}

func sumProducts(ch chan float64, products ...Products) (float64, error) {
	// price * count
	var totalPrice float64

	for _, proproduct := range products {
		if proproduct.Price < 0 {
			return 0, errors.New("Price invalid")
		}
		fmt.Println(proproduct.Price)
		totalPrice += totalPrice + proproduct.Price + 1
	}
	ch <- totalPrice
	return totalPrice, nil
}

func calcHalfHours(mins float64) float64 {
	return mins / 30
}

func sumServices(ch chan float64, services ...Services) (float64, error) {
	// precio * media hora trabajada, si no llega a trabajar 30
	// minutos se le cobra como si hubiese trabajado media hora
	var totalPrice float64
	var totalMins float64
	for _, service := range services {
		if service.Price < 0 {
			return 0, errors.New("Price invalid")
		}
		//fmt.Println(service.Price)
		totalMins += totalMins + service.MinWorked
	}
	ch <- totalPrice
	return totalPrice, nil
}

func sumMaintenance(ch chan float64, maintenances ...Maintenance) (float64, error) {
	// price * count
	var totalPrice float64

	for _, maintenance := range maintenances {
		if maintenance.Price < 0 {
			return 0, errors.New("Price invalid")
		}
		//fmt.Println(maintenance.Price)
		totalPrice += totalPrice + maintenance.Price
	}
	ch <- totalPrice
	return totalPrice, nil
}

func main() {

	products := [3]Products{
		{
			Name:  "p1",
			Price: 1,
			count: 1,
		},
		{
			Name:  "p2",
			Price: 2,
			count: 2,
		},
		{
			Name:  "p3",
			Price: 3,
			count: 3,
		},
	}
	services := [3]Services{
		{
			Name:      "s1",
			Price:     1,
			MinWorked: 140,
		},
		{
			Name:      "s2",
			Price:     2,
			MinWorked: 150,
		},
		{
			Name:      "s3",
			Price:     3,
			MinWorked: 160,
		},
	}
	maintenances := [3]Maintenance{
		{
			Name:  "m1",
			Price: 1,
		},
		{
			Name:  "m2",
			Price: 2,
		},
		{
			Name:  "m3",
			Price: 3,
		},
	}

	productChanel := make(chan float64)
	servicesChanel := make(chan float64)
	maintenaceChanel := make(chan float64)

	for i := 0; i <= 2; i++ {
		go sumProducts(productChanel, products[i])
		go sumServices(servicesChanel, services[i])
		go sumMaintenance(maintenaceChanel, maintenances[i])
	}
	var productReading float64
	var servideReading float64
	var maintenaceReading float64
	for i := 0; i <= 2; i++ {
		productReading = <-productChanel
		//fmt.Println("Reading  product chanel ", productReading)
		servideReading = <-servicesChanel
		//fmt.Println("Reading  service chanel ", servideReading)
		maintenaceReading = <-maintenaceChanel
		//fmt.Println("Reading  maintenance chanel ", maintenaceReading)
	}
	fmt.Println("Total price for PRODUCTS: ", productReading)
	fmt.Println("Total price for SERVICES: ", servideReading)
	fmt.Println("Total price for MAINTENACE: ", maintenaceReading)

}
