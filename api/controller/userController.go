package controller

import (
	"api/auth"
	"api/database"
	"api/models"

	"fmt"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	Name string
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

type UserGetResponse struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}

func CreateUser(c *gin.Context) {
	req := &UserCreateRequest{}
	err := c.Bind(req)
	if err != nil {
		fmt.Println(err)
	}

	var token string
	token = auth.GenerateToken(req.Name)

	user := &models.User{Name: req.Name, Token: token}
	err = user.CreateUser(database.GetDB())
	if err != nil {
		fmt.Println(err)
	}

	res := &UserCreateResponse{Token: token}

	c.JSON(200, res)
}

func GetUser(c *gin.Context) {
	token := c.GetHeader("x-token")

	user, err := models.GetUser(database.GetDB(), token)
	if err != nil {
		fmt.Println(err)
	}

	res := &UserGetResponse{Name: user.Name}

	c.JSON(200, res)
}

func UpdateUser(c *gin.Context) {
	token := c.GetHeader("x-token")
	req := &UserUpdateRequest{}
	err := c.Bind(req)
	if err != nil {
		fmt.Println(err)
	}

	user, err := models.GetUser(database.GetDB(), token)
	if err != nil {
		fmt.Println(err)
	}

	err = user.UpdateUser(database.GetDB(), req.Name)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{})
}
