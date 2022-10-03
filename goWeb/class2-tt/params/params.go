package main

import "github.com/gin-gonic/gin"

// Definimos una pseudobase de datos donde consultaremos la información.
var empleados = map[string]string{
	"644": "Empleado A",
	"755": "Empleado B",
	"777": "Empleado C",
}

func main() {
	server := gin.Default()
	server.GET("/", Home)
	server.GET("/empleados/:id", GetOne)
	server.Run(":8080")
}

// Este handler se encargará de responder a /.
func Home(ctx *gin.Context) {
	ctx.String(200, "¡Bienvenido a la Empresa Gophers!")
}

// Este handler verificará si la id que pasa el cliente existe en nuestra base de datos.
func GetOne(ctx *gin.Context) {
	empleado, err := empleados[ctx.Param("id")]
	if err {
		ctx.String(200, "Información del empleado %s, nombre: %s", ctx.Param("id"), empleado)
	} else {
		ctx.String(404, "Información del empleado ¡No existe!")
	}
}
