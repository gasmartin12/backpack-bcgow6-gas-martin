package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class3-tm/pkg/web"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		c.AbortWithStatusJSON(web.NewResponse(http.StatusUnauthorized, nil, "does not have permissions to perform the requested request"))
		return
	}
	c.Next()
}

func IdValidationMiddleware(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(web.NewResponse(http.StatusBadRequest, nil, "invalid ID"))
		return
	}
	c.Next()
}
