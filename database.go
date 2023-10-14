package database

import (
	"fmt"
	"go-auth/config"
	"go-auth/models"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(192.168.1.103:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", config.Config("DB_USERNAME"), config.Config("DB_PASSWORD"), port, config.Config("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Ava{})
	DB.AutoMigrate(&models.Photo{})
	DB.AutoMigrate(&models.Video{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Post{})
	fmt.Println("Database Migrated")

}
