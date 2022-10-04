/*
Ejercicio 2 - Hola {nombre}

Crea dentro de la carpeta go-web un archivo llamado main.go
Crea un servidor web con Gin que te responda un JSON que tenga una
clave "message" y diga Hola seguido por tu nombre.
Pegale al endpoint para corroborar que la respuesta sea la correcta.
*/
/*
Ejercicio 3 - Listar Entidad

Ya habiendo creado y probado nuestra API que nos saluda, generamos una
ruta que devuelve un listado de la temática elegida.
Dentro del "main.go", crea una estructura según la temática con los campos
correspondientes.
Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo:
"/productos"
Genera un handler para el endpoint llamado "GetAll".
Crea una slice de la estructura, luego devuelvela a través de nuestro
endpoint.
*/

package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Lastname     string  `json:"lastname"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float64 `json:"height"`
	Active       bool    `json:"isActive"`
	CreationDate string  `json:"createdAt"`
}

func main() {
	router := gin.Default()
	// exer 2
	name := "Gas"
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello 👋" + name,
		})
	})

	// ejer 3
	//----------------------------- from Code
	router.GET("/users", GetAllCode)
	//----------------------------- from JSON
	router.GET("/users1", GetAllJson)

	router.Run()

}

// ----------------------------- code
func generateUsers() []User {
	users := []User{
		{1, "test1", "test1", "test1@mercadolibre.com", 34, 1.76, true, "3-10-2022"},
		{2, "test2", "test2", "test2@mercadolibre.com", 12, 1.67, false, "2-10-2022"},
		{3, "test3", "test3", "test3@mercadolibre.com", 46, 1.60, true, "1-10-2022"},
	}
	return users
}

func GetAllCode(context *gin.Context) {
	context.JSON(http.StatusOK, generateUsers())
}

// ----------------------------- json

func GetUsers() ([]User, error) {
	var users []User
	raw, err := ioutil.ReadFile("./go-web/users.json")
	if err != nil {
		return nil, errors.New("error when read the file")
	}
	json.Unmarshal(raw, &users)
	return users, nil
}

func GetAllJson(context *gin.Context) {
	products, err := GetUsers()
	if err == nil {
		context.JSON(http.StatusOK, products)
	}
}
