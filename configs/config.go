package configs

import (
	"alterra/models/users"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	dsn := "root:Smkn1grati.@tcp(127.0.0.1:3307)/pesanbuku?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("DB failed connect")
	} else {
		fmt.Println("Connection Estabilished")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&users.User{})
}
