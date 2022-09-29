package main

import (
	"fmt"
)

var SalaryMin = 100000

func main() {

	var age int
	var isEmployee string
	var seniority int
	var salary int

	fmt.Print("Insert your age : ")
	fmt.Scanf("%d", &age)

	fmt.Print("Is a employee (y/n): ")
	fmt.Scanf("%s", &isEmployee)

	fmt.Print("Insert yours years of seniority : ")
	fmt.Scanf("%d", &seniority)

	fmt.Print("Insert yours salary : ")
	fmt.Scanf("%d", &salary)

	if age < 22 {
		fmt.Println("To receive the loan you must be over 22 years old.")
	} else if isEmployee == "n" {
		fmt.Println("You must be an employee to receive the loan.")
	} else if seniority < 1 {
		fmt.Println("Your seniority must be more than one year to receive the loan.")
	} else if salary < SalaryMin {
		fmt.Println("Congratulations, you can receive a loan with interest!")
	} else {
		fmt.Println("Congratulations, you can receive a interest-free loan !")
	}

}
