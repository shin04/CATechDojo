package main

import (
	"api/config"
	"api/database"
	"api/router"
)

func main() {
	config := &config.Config{}
	config.Init()

	database.Init(config)

	router := &router.Router{}
	router.Init()
	router.Run()
}
