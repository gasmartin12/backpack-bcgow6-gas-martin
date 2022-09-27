package main

// capa de abstraccion muy alta para trabajar con el sistema Operativo
// conocer info de un archivo, trabajar con variables de entorno del sistema

import (
	"fmt"
	"os"
)

func main() {

	// os.Setenv()
	err := os.Setenv("NAME", "Gaston")
	if err != nil {
		panic(err)
	}

	// os.Getenv()
	name := os.Getenv("NAME")
	if name != "" {
		fmt.Printf(name)
	} else {
		fmt.Printf("Empty variable")
	}

	// os.LookupEnv()
	dir, ok := os.LookupEnv("HOME")
	if !ok {
		panic("The variable HOME dosent exist")
	}
	fmt.Println("ENV : ", dir)

	// os.Exit()
	flag := true
	if flag {
		fmt.Println("flag is true")
	} else {
		fmt.Println("flag is false")
		os.Exit(1)
	}
	fmt.Println("End instruction")

	//os.ReadDir()
	entries, err := os.ReadDir("..")
	if err != nil {
		fmt.Printf("Collected entries: %s\n", entries)
		panic(err)
	}
	for _, value := range entries {
		fmt.Println(value.Name())
	}

	//os.ReadFile()
	fileContectAsByte, err := os.ReadFile("../into/go.mod")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Content: %s\n", fileContectAsByte)

	//os.WriteFile()
	message := "Hello world! i'm in a file :)"
	os.WriteFile(
		"./text.txt",
		[]byte(message),
		0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("File created succesfully")

}
