package main

import "fmt"

type Student struct {
	Name     string
	Lastname string
	DNI      string
	Date     string
}

func main() {
	student := Student{
		"Gas", "Martin", "41522639", "12/03/1999",
	}
	details(student)

	student.detailsMethod()
}

func details(student Student) {
	fmt.Printf("Detalles: %+v\n", student)
}

func (student Student) detailsMethod() {
	fmt.Printf("Name: %s %s\n", student.Name, student.Lastname)
	fmt.Printf("DNI: %s\n", student.DNI)
	fmt.Printf("Fecha: %s\n", student.Date)
}
