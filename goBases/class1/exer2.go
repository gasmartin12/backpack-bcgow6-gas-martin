package main

import "fmt"

func main() {
	humidity := 10                       //int
	temperature := 16.3                  //float64
	atmospheric_pressure := 5            //int
	city := "Rio cuarto - Cordoba - Arg" //string

	fmt.Println("The temperature for ", city, " is: ", temperature)
	fmt.Println("The humidity for ", city, " is: ", humidity)
	fmt.Println("The atmospheric pressure for ", city, " is: ", atmospheric_pressure)
}
