package main

import "fmt"

func main() {
	months := map[int]string{1: "junuary", 2: "February",
		3: "March", 4: "April", 5: "May", 6: "June", 7: "July",
		8: "August", 9: "September", 10: "October", 11: "November",
		12: "December"}

	var numMonth int

	fmt.Print("Insert a month (1..12): ")
	fmt.Scanf("%d", &numMonth)

	x, ok := months[numMonth]
	if !ok {
		fmt.Println("Invalid month")
	}
	fmt.Println(x)

}
