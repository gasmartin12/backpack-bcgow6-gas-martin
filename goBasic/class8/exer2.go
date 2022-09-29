/*
Ejercicio 2 - Registrando clientes

El mismo estudio del ejercicio anterior, solicita una funcionalidad para
poder registrar datos de nuevos clientes. Los datos requeridos para
registrar a un cliente son:
Legajo
Nombre y Apellido
DNI
Número de teléfono
Domicilio

Tarea 1: El número de legajo debe ser asignado o generado por separado y
en forma previa a la carga de los restantes gastos. Desarrolla e implementa
una función para generar un ID que luego utilizarás para asignarlo como
valor a "Legajo". Si por algún motivo esta función retorna valor "nil",
debe generar un panic que interrumpa la ejecución y aborte.
Tarea 2: Antes de registrar a un cliente, debes verificar si el mismo ya
existe. Para ello, necesitas leer los datos de un archivo .txt. En algún
lugar de tu código, implementa la función para leer un archivo llamado
"customers.txt" (como en el ejercicio anterior, este archivo no existe,

	por lo que la función que intente leerlo devolverá un error).
	Debes manipular adecuadamente ese error como hemos visto hasta aquí.
	Ese error deberá:

1.- generar un panic;
2.- lanzar por consola el mensaje: "error: el archivo indicado no fue
encontrado o está dañado", y continuar con la ejecución del programa
normalmente.
Tarea 3: Luego de intentar verificar si el cliente a registrar ya existe,
desarrolla una función para validar que todos los datos a registrar de un
cliente contienen un valor distinto de cero. Esta función debe retornar,
al menos, dos valores. Uno de los valores retornados deberá ser de tipo error
para el caso de que se ingrese por parámetro algún valor cero
(recuerda los valores cero de cada tipo de dato, ej: 0, ", nil).
Tarea 4: Antes de finalizar la ejecución, incluso si surgen panics,

	se deberán imprimir por consola los siguientes mensajes: "Fin de la ejecución",
	 "Se detectaron varios errores en tiempo de ejecución" y
	 "No han quedado archivos abiertos" (en ese orden).
	 Utiliza defer para cumplir con este requerimiento.

Requerimientos generales:
Utiliza recover para recuperar el valor de los panics que puedan surgir
(excepto en la tarea 1).
Recordá realizar las validaciones necesarias para cada retorno que
pueda contener un valor error (por ejemplo las que intenten leer archivos).
Genera algún error, personalizandolo a tu gusto, utilizando alguna de las
funciones que GO provee para ello (realiza también la validación pertinente
para el caso de error retornado).
*/
package main

import (
	"errors"
	"fmt"
	"os"
)

type Client struct {
	File        int
	FirstName   string
	LastName    string
	Id          int
	PhoneNunber int
	Address     string
}

type NameError struct {
	msg string
}

func (ne *NameError) Error() string {
	return ne.msg
}

func idGenerator() func(string) (int, error) {
	id := 0
	return func(name string) (int, error) {
		if len(name) < 2 {
			return 0, &NameError{"Name is too short"}
		}
		id++
		return id, nil
	}
}

func openFile(name string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	file, err := os.ReadFile(name)
	if err != nil {
		panic("File not found or is corrupted")
	}
	if len(file) > 0 {
		fmt.Println("File found")
	}
}

func validateClient(client Client) (bool, error) {
	switch {
	case client.File == 0:
		return false, errors.New("error: File empty")
	case client.FirstName == "":
		return false, errors.New("error: First name empty")
	case client.LastName == "":
		return false, errors.New("error: Last name empty")
	case client.Id == 0:
		return false, errors.New("error: Id empty")
	case client.PhoneNunber == 0:
		return false, errors.New("error: PhoneNunber empty")
	case client.Address == "":
		return false, errors.New("error: Address empty")
	}
	return true, nil
}

func main() {
	defer func() {
		err := recover()
		fmt.Println("Process Done")
		if err != nil {
			fmt.Println("runtime error")
			fmt.Println("files closed")
		}
	}()

	genId := idGenerator()

	client := new(Client)

	client.FirstName = "Gas"
	id, err1 := genId(client.FirstName)

	if err1 != nil {
		panic(err1)
	}

	openFile("./customers.txt.")
	fmt.Println("Gas's id is ", id)

	isOk, err2 := validateClient(*client)
	if err2 != nil {
		panic("Fields are missing")
	}
	if isOk {
		fmt.Println("Registering client")
	}
}
