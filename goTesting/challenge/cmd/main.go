package main

import (
	"log"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/challenge/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	if err := r.Run(":18085"); err != nil {
		log.Fatalf("Error running engine")
	}

}
