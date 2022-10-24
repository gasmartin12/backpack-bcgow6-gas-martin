package main

import (
	"clase2_parte1/directorio"
	"fmt"
)

func main() {
	db := directorio.NewDB()
	motor := directorio.NewEngine(db)

	fmt.Printf("La versión actual es %s\n", motor.GetVersion())
}
