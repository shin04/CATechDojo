package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB struct {
		Host     string
		Name     string
		Port     string
		Username string
		Password string
	}
}

func (config *Config) Init() {
	var envfilepath string

	env := os.Getenv("GO_ENV")
	if env == "" {
		envfilepath = "../.env"
	} else if env == "test" {
		envfilepath = "../../.env.test"
	}

	err := godotenv.Load(envfilepath)
	if err != nil {
		fmt.Println(err)
		println("Error loading .env file")
	}
	config.DB.Host = os.Getenv("DB_HOST")
	config.DB.Name = os.Getenv("DB_NAME")
	config.DB.Port = os.Getenv("DB_PORT")
	config.DB.Username = os.Getenv("DB_USER")
	config.DB.Password = os.Getenv("DB_PASS")
}
