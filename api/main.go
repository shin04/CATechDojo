package main

import (
	"api/database"
	"api/router"
)

func main() {
	database.Init()

	router.Init()
	router := router.GetRouter()
	router.Run()
}
