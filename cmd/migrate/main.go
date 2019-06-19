package main

import (
	"fmt"
	"log"
	"os"
	"github.com/whiteblock/dexter"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connect := fmt.Sprintf("host=%s user=%s dbname=%s password=%s", "localhost", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	db, err := gorm.Open("postgres", connect)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}
	defer db.Close()
	db.AutoMigrate(&dexter.Chart{})
	db.AutoMigrate(&dexter.Alert{})
	db.AutoMigrate(&dexter.Webhook{})
	db.AutoMigrate(&dexter.Indicator{})
}
