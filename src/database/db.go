package database

import (
	"log"
	"time"

	"GoPacientes/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func StartDB() {

	dsn := "root:@tcp(127.0.0.1:3306)/myteste?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	database.AutoMigrate(models.Paciente{})

	db = database

	config, _ := db.DB()
	config.SetConnMaxLifetime(time.Hour)
	config.SetMaxOpenConns(100)
	config.SetMaxIdleConns(10)

}

func GetDB() *gorm.DB {
	return db

}
