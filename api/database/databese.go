package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Init initializes database
func Init() {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		println("Error loading .env file")
	}
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":3306)/" + DB_NAME
	dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		println("db connection success !!")
	}
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}
