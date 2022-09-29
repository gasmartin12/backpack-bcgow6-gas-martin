package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.ReadFile("file not found")
	if err != nil {
		panic(err)
	}
	fmt.Println("All right")
}
