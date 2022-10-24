package main

import (
	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/challenge/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	r.Run(":18085")

}
