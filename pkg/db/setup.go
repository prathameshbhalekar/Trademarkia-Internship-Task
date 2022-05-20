package db

import (
	"log"

	"github.com/prathameshbhalekar/trademarkia-task/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(url string) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Users{}, &models.Like{})

	DB = db
}
