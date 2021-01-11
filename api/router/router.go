package router

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello!!!")
	})

	// user api
	router.POST("/user/create", controller.CreateUser)
	router.GET("/user/get", controller.GetUser)
	router.PUT("/user/update", controller.UpdateUser)
}

func GetRouter() *gin.Engine {
	return router
}
