/*
Ejercicio 1 - Filtremos nuestro endpoint
Según la temática elegida, necesitamos agregarles filtros a nuestro
endpoint, el mismo se tiene que poder filtrar por todos los campos.
Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
Luego genera la lógica de filtrado de nuestro array.
Devolver por el endpoint el array filtrado.

Ejercicio 2 - Get one endpoint
Generar un nuevo endpoint que nos permita traer un solo resultado del
array de la temática. Utilizando path parameters el endpoint debería ser
/temática/:id (recuerda que siempre tiene que ser en plural la temática).
Una vez recibido el id devuelve la posición correspondiente.
Genera una nueva ruta.
Genera un handler para la ruta creada.
Dentro del handler busca el item que necesitas.
Devuelve el item según el id.
Si no encontraste ningún elemento con ese id devolver como código de
respuesta 404.

*/

package main

import (
	"fmt"
	"net/http"
	"strconv"

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

	router.GET("/fillUser", FillUserHandler)
	router.GET("/users/:id", GetUserByIdHander)
	router.Run()

}

// exer 1
func FillUserHandler(c *gin.Context) {
	usersList := createUsers()

	var user User
	if c.ShouldBindQuery(&user) == nil {
		fmt.Println(user.Id)
		fmt.Println(user.Name)
		fmt.Println(user.Lastname)
		fmt.Println(user.Email)
		fmt.Println(user.Age)
		fmt.Println(user.Height)
		fmt.Println(user.Active)
		fmt.Println(user.CreationDate)
	}

	var userFilter []*User

	for _, u := range usersList {
		if user.Id == u.Id && user.Name == u.Name && user.Lastname == u.Lastname && user.Email == u.Email && user.Age == u.Age && user.Height == u.Height && user.Active == u.Active && user.CreationDate == u.CreationDate {
			userFilter = append(userFilter, u)
		}
	}
	c.JSON(http.StatusOK, userFilter)
}

// exer 2
func GetUserByIdHander(c *gin.Context) {
	usersList := createUsers()
	var user User

	id, _ := strconv.Atoi(c.Param("id"))
	isExist := false

	for _, u := range usersList {
		if u.Id == int(id) {
			isExist = true
			user = *u
			break
		}
	}

	if !isExist {
		//c.JSON(404, nil)
		c.JSON(http.StatusNotFound, nil)
	} else {
		//c.JSON(200, nil)
		c.JSON(http.StatusOK, user)
	}
}

// modularizando para usar en los 2 ejercicios
func createUsers() []*User {
	usersList := []*User{
		{1, "test1", "test1", "test1@mercadolibre.com", 34, 1.76, true, "3-10-2022"},
		{2, "test2", "test2", "test2@mercadolibre.com", 12, 1.67, false, "2-10-2022"},
		{3, "test3", "test3", "test3@mercadolibre.com", 46, 1.60, true, "1-10-2022"},
	}
	return usersList
}
