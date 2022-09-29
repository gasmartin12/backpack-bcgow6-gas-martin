package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	var employeeName string
	fmt.Println("Employees list")
	fmt.Println(employees)

	fmt.Print("Insert the name of the employee: ")
	fmt.Scanf("%s", &employeeName)

	x, ok := employees[employeeName]
	if !ok {
		fmt.Println("This employee doesn't exist")
	}
	fmt.Println("The age is: ", x)

	count := 0
	for name, age := range employees {
		if age > 21 {
			count++
			fmt.Println(name)
		}
	}
	fmt.Println("Employees over 21 years of age: ", count)

	var addEmployee string
	fmt.Print("Add Federico as employee? (y/n): ")
	fmt.Scanf("%s", &addEmployee)
	if addEmployee == "y" {
		employees["Federico"] = 25
	}
	fmt.Println("Updated list")
	fmt.Println(employees)

	var deleteEmployee string
	fmt.Print("Delete Pedro from the list? (y/n): ")
	fmt.Scanf("%s", &deleteEmployee)
	if deleteEmployee == "y" {
		delete(employees, "Pedro")
	}
	fmt.Println("Updated list")
	fmt.Println(employees)
}
