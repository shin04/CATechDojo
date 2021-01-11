package router

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

// var router *gin.Engine

type Router struct {
	Engin *gin.Engine
}

func (router *Router) Init() {
	router.Engin = gin.Default()

	router.Engin.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello!!!")
	})

	// user api
	router.Engin.POST("/user/create", controller.CreateUser)
	router.Engin.GET("/user/get", controller.GetUser)
	router.Engin.PUT("/user/update", controller.UpdateUser)
}

func (router *Router) Run() {
	router.Engin.Run()
}
