package database

import (
	"fmt"
	"log"
	"strateegy/project-service/database/migrations"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	strconn := "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=9932"

	database, err := gorm.Open(postgres.Open(strconn), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDB() *gorm.DB {
	return db
}
