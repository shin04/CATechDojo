package router

import (
	"api/controller"
	"api/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// var router *gin.Engine

type Router struct {
	Engin *gin.Engine
}

func (router *Router) Init() {
	router.Engin = gin.Default()

	router.Engin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"content-type", "Origin", "x-token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Engin.Use(middleware.ErrorHandler())

	router.Engin.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello!!!")
	})

	// user api
	router.Engin.GET("/users", controller.GetAllUser)
	router.Engin.POST("/user/create", controller.CreateUser)
	router.Engin.GET("/user/get", controller.GetUser)
	router.Engin.PUT("/user/update", controller.UpdateUser)
}

func (router *Router) Run() {
	router.Engin.Run()
}
