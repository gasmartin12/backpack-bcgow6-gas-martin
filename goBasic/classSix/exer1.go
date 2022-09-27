/*
Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una
estructura usuario con funciones que vayan agregando
información a la estructura. Para optimizar y ahorrar
memoria requieren que la estructura de usuarios ocupe
el mismo lugar en memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad,
Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.
*/

package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name    string
	Surname string
	Age     int
	Email   string
	Pass    string
}

func updateName(name *string, surname *string) error {
	*name = "Gaston"
	*surname = "Mar"
	if *name == "" || *surname == "" {
		return nil
	}
	return errors.New("Invalid name or surname")
}
func updateAge(age *int) error {
	*age = 24
	if *age < 0 && *age >= 100 {
		return nil
	}
	return errors.New("Invalid age")

}
func updateEmail(email *string) error {
	*email = "gaastonmaartin12@gmail.com"
	if *email == "" {
		return nil
	}
	return errors.New("Invalid email")
}
func updatePass(pass *string) error {
	*pass = "password"
	if *pass == "" {
		return nil
	}
	return errors.New("Invalid password")
}

func printUser(u User) {
	fmt.Println("---------------")
	fmt.Printf("Hi %s %s!\nAge: %d\nEmail: %s\nPass: %s\n",
		u.Name, u.Surname, u.Age, u.Email, u.Pass)
	fmt.Println("---------------")
}
func isNameOk(u User) {
	err := updateName(&u.Name, &u.Surname)
	if err != nil {
		fmt.Printf("New name: %s\nNew surname: %s\n", u.Name, u.Surname)
	}
}
func isAgeOk(u User) {
	err := updateAge(&u.Age)
	if err != nil {
		fmt.Printf("New age: %d\n", u.Age)
	} else {
		fmt.Print("asd")
	}
}
func isEmailOk(u User) {
	err := updateEmail(&u.Email)
	if err != nil {
		fmt.Printf("New email: %s\n", u.Email)
	}
}
func isPassOk(u User) {
	err := updatePass(&u.Pass)
	if err != nil {
		fmt.Printf("New pass: %s\n", u.Pass)
	}
}

func main() {
	u := User{
		Name:    "Gas",
		Surname: "Martin",
		Age:     23,
		Email:   "gaston.gmartin@mercadolibre.com",
		Pass:    "Pass",
	}
	printUser(u)
	isNameOk(u)
	isAgeOk(u)
	isEmailOk(u)
	isPassOk(u)
}
