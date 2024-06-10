package db

import (
	"all_inclusive_app/config"
	"all_inclusive_app/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB инициализирует базу данных
func InitDB() *gorm.DB {

	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",	
			config.ConfigDB.Host, config.ConfigDB.User, config.ConfigDB.Password, config.ConfigDB.DBName, config.ConfigDB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&models.User{})

	return db
}

func GetFromDB() *gorm.DB {
	return db
}