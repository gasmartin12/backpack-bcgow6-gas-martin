package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/bootcampers", Bootcampers)

	fmt.Println("Iniciando el servidor . . . 🚀 en el puerto 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func Bootcampers(w http.ResponseWriter, r *http.Request) {
	type Message struct {
		Message string
	}

	m := Message{Message: "👋 Hola Bootcampers"}
	data, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(data)
}
