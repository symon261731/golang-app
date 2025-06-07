package main

import (
	"github.com/joho/godotenv"
	"golang-app/internal/link"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")))
	if err != nil {
		panic(err)
	}

	migrateError := db.AutoMigrate(&link.Link{})

	if migrateError != nil {
		panic(err)
	}

	log.Print("Successfully migrated database")
}
