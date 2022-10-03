package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func httpExample() {

	type product struct {
		Name      string
		Price     float64
		Published bool
	}

	jsonData := []byte(`{"Name": "MacBook Air", "Price": 900, "Published": true}`)

	var p product
	if err := json.Unmarshal(jsonData, &p); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product: ", p)

	http.HandleFunc("/hola", holaHandler)
	http.ListenAndServe(":8080", nil)

}

func holaHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hola\n")
}
