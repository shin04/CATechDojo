package database

import (
	"api/config"

	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init(config *config.Config) {
	DB_USER := config.DB.Username
	DB_PASS := config.DB.Password
	DB_HOST := config.DB.Host
	DB_NAME := config.DB.Name
	DB_PORT := config.DB.Port

	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME
	dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		println("db connection success !!")
	}
}

func GetDB() *gorm.DB {
	return db
}
