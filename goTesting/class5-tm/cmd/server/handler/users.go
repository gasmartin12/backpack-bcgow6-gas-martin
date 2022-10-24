package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class5-tm/internal/users"
	"github.com/gasmartin12/backpack-bcgow6-gas-martin/goTesting/class5-tm/pkg/web"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Height   int    `json:"height" binding:"required"`
}

type RequestUpdate struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Height   int    `json:"height" binding:"required"`
	Active   *bool  `json:"active" binding:"required"`
}

type RequestUpdateLastNameAndAge struct {
	LastName string `json:"lastName" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(service users.Service) *User {
	return &User{service: service}
}

// ListUser godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 404 {object} web.Response
// @Router /users [get]
func (u *User) GetAll(c *gin.Context) {
	users, err := u.service.GetAll()
	if err != nil {
		c.JSON(web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, users, ""))
}

// StoreUser godoc
// @Summary Store user
// @Tags Users
// @Description store user
// @accept json
// @Produce  json
// @Param token header string true "token"
// @Param product body Request true "User to store"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 400 {object} web.Response
// @failure 500 {object} web.Response
// @Router /users [post]
func (u *User) Store(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	user, err := u.service.Store(request.Name, request.LastName, request.Email, request.Age, request.Height)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, user, ""))
}

// UpdateUser godoc
// @Summary Update user
// @Tags Users
// @Description update user
// @accept json
// @Produce  json
// @Param token header string true "token"
// @Param id path string true "User id"
// @Param product body RequestUpdate true "User to update"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 400 {object} web.Response
// @failure 404 {object} web.Response
// @failure 500 {object} web.Response
// @Router /users/{id} [put]
func (u *User) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var request RequestUpdate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	user, err := u.service.Update(id, request.Name, request.LastName, request.Email, request.Age, request.Height, *request.Active)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, user, ""))
}

// UpdateLastnameAgeUser godoc
// @Summary Update lastname and age user
// @Tags Users
// @Description update lastname and age user
// @accept json
// @Produce  json
// @Param token header string true "token"
// @Param id path string true "User id"
// @Param product body RequestUpdateLastNameAndAge true "lastname and age to update"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 400 {object} web.Response
// @failure 404 {object} web.Response
// @failure 500 {object} web.Response
// @Router /users/{id} [patch]
func (u *User) UpdateLastNameAndAge(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var request RequestUpdateLastNameAndAge
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	user, err := u.service.UpdateLastNameAndAge(id, request.LastName, request.Age)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, user, ""))
}

// DeleteUser godoc
// @Summary Delete user
// @Tags Users
// @Description delete user
// @Produce  json
// @Param token header string true "token"
// @Param id path string true "User id"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 400 {object} web.Response
// @failure 404 {object} web.Response
// @Router /users/{id} [delete]
func (u *User) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := u.service.Delete(id)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, gin.H{"message": fmt.Sprintf("the user with id %d was deleted", id)}, ""))
}
