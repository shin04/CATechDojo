package main

import (
	"api/database"
	"api/router"
)

func main() {
	database.Init()

	router := &router.Router{}
	router.Init()
	router.Run()
}
