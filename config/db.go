package config

import (
	"final-project/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	dbPort   = "5432"
	dbName   = "go-final-project"
	user     = "postgres"
	password = "root"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("gagal terhubung ke Database :", err)
	}

	defer fmt.Println("Sukses terhubung ke Database")

	db.Debug().AutoMigrate(entity.User{}, entity.Comment{}, entity.Photo{}, entity.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
