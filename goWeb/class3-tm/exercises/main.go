/*
Ejercicio 1 - Crear Entidad
Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
siguientes pasos:
1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global).

Ejercicio 2 - Validación de campos
Se debe implementar las validaciones de los campos al momento de enviar la petición, para
eso se deben seguir los siguientes pasos:
1. Se debe validar todos los campos enviados en la petición, todos los campos son
requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400
con el mensaje "el campo %s es requerido".
(En %s debe ir el nombre del campo que no está completo).

Ejercicio 3 - Validar Token
Para agregar seguridad a la aplicación se debe enviar la petición con un token, para eso se
deben seguir los siguientes pasos::

1. Al momento de enviar la petición se debe validar que un token sea enviado
2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
mensaje que "no tiene permisos para realizar la petición solicitada".
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

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

const (
	filePath = "./users.json"
	Token    = "123Pass"
)

var users []User

func main() {

	fmt.Println(ReadJson(filePath))

	router := gin.Default()

	user, err := ReadJson(filePath)
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	users = user

	userRouter := router.Group("/users")
	{
		userRouter.POST("/", AddUser())
	}

	router.Run()
}

func AddUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != Token {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "invalid token"})
			return
		}

		var u User
		if err := ctx.ShouldBindJSON(&u); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("field %s is required", err.Error())})
			return
		}
		u.Id = len(users) + 1
		users = append(users, u)
		ctx.JSON(http.StatusOK, gin.H{"User added": u})
	}

}

func ReadJson(filePath string) ([]User, error) {

	var transactions []User

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("error reading the file")
	}
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	return transactions, nil

}
