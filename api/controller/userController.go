package controller

import (
	"api/auth"
	"api/database"
	"api/models"

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
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	var token string
	token = auth.GenerateToken(req.Name)

	user := &models.User{Name: req.Name, Token: token}
	err = user.CreateUser(database.GetDB())
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	res := &UserCreateResponse{Token: token}

	c.JSON(200, res)
}

func GetAllUser(c *gin.Context) {
	users, err := models.GetAllUser(database.GetDB())
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	token := c.GetHeader("x-token")

	user, err := models.GetUser(database.GetDB(), token)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	res := &UserGetResponse{Name: user.Name}

	c.JSON(200, res)
}

func UpdateUser(c *gin.Context) {
	token := c.GetHeader("x-token")
	req := &UserUpdateRequest{}
	err := c.Bind(req)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	user, err := models.GetUser(database.GetDB(), token)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	err = user.UpdateUser(database.GetDB(), req.Name)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(200, gin.H{})
}
