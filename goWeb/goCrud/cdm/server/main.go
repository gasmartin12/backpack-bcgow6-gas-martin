package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goWeb/goCrud/cmd/server/handler"
	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goWeb/goCrud/internal/users"
	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goWeb/goCrud/pkg/store"
	//"github.com/ncostamagna/meli-bootcamp/docs" -> paquete generado con swag init -g cdm/server/main.go
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the .env file")
	}

	db := store.New(store.FileType, "./users.json")

	repo := users.NewRepository(db)
	service := users.NewService(repo)

	u := handler.NewUser(service)

	router := gin.Default()

	routerGroup := router.Group("/users")
	routerGroup.POST("/", u.SaveUser())
	routerGroup.GET("/", u.GetAllUSer())
	routerGroup.PUT("/:id", u.UpdateUser())
	routerGroup.DELETE("delete/:id", u.DeleteUser())
	routerGroup.PATCH("updateName/:id", u.Update())
	router.Run()
}
