package main

func main() {
	server := gin.Default()
	server.GET("/", HandlerRaiz)
	//Ahora podemos atender peticiones a /gophers/, /gophers/get o /gophers/info de una
	//forma más elegante.
	gopher := server.Group("/gophers")
	{
		gopher.GET("/", HandlerGophers)
		gopher.GET("/get", HandlerGetGopher)
		gopher.GET("/info", HandlerGetInfo)
	}
	server.GET("/about", HandlerAbout)
	server.Run(":8080")
}
