package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goWeb/goCrud/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type request struct {
	Id         int     `json:"Id"`
	Name       string  `form:"Name" json:"Name" validate:"required"`
	LastName   string  `form:"LastName" json:"LastName" validate:"required"`
	Email      string  `form:"Email" json:"Email" validate:"required"`
	Age        int     `form:"Age" json:"Age" validate:"required"`
	Height     float64 `form:"Height" json:"Height" validate:"required"`
	Active     *bool   `form:"Active" json:"Active" validate:"required"`
	DateCreate string  `form:"DateCreate" json:"DateCreate" validate:"required"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (r *request) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

/*
// ListUser godoc
// @Summary List users
// @Tags Products
// @Description get users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
*/
func (u *User) GetAllUsers(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") || token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	users, err := u.service.GetAllUsers()
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if len(users) <= 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"error": "No hay usuarios cargados"})
		return
	} else {
		ctx.JSON(200, users)
	}
}

func (u *User) SaveUser(c *gin.Context) {

	token := c.GetHeader("token")
	if token != "password" || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		return
	}

	var req request

	c.ShouldBindJSON(&req)
	if err := req.validate(); err != nil {
		valError := err.(validator.ValidationErrors)
		msg := []string{}
		for _, e := range valError {
			msg = append(msg, "el campo "+e.Field()+" es requerido")
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	userSave, err := u.service.SaveUser(req.Name, req.LastName, req.Email, req.Age, req.Height, *req.Active, req.DateCreate)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userSave)
}

func (u *User) UpdateUser(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "password" || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(404, gin.H{"error": "Id inválido"})
		return
	}

	var req request

	c.ShouldBindJSON(&req)
	if err := req.validate(); err != nil {
		valError := err.(validator.ValidationErrors)
		msg := []string{}
		for _, e := range valError {
			msg = append(msg, "el campo "+e.Field()+" es requerido")
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	userUpdate, err := u.service.UpdateUser(int(id), req.Name, req.LastName, req.Email, req.Age, req.Height, *req.Active, req.DateCreate)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, userUpdate)
}

/*
// Delete User
// @Sumary Delete User
// @Tags User
// @Param id path in true "User ID"
// @Param token header string true "Token"
// @Success 204 
// @Failture 400 (object) web.Response
// @Failture 404 (object) web.Response
// @Router /users/{id} [DELETE]
*/
func (u *User) DeleteUser(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "password" || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(404, gin.H{"error": "Id inválido"})
		return
	}

	err = u.service.DeleteUser(int(id))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": fmt.Sprintf("El usuario %d ha sido eliminado correctamente", id)})
}

func (u *User) UpdatePatch(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "password" || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(404, gin.H{"error": "Id inválido"})
		return
	}

	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if req.LastName == "" {
		c.JSON(400, gin.H{"error": "El campo Apellido no puede estar vacío"})
		return
	}
	if req.Age <= 0 {
		c.JSON(400, gin.H{"error": "El campo Edad no puede ser menor a 1"})
		return
	}

	userUpdate, err := u.service.UpdatePatch(int(id), req.LastName, req.Age)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, userUpdate)
}
