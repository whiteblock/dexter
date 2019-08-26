package main

import (
	"log"
	"os"
	"github.com/whiteblock/dexter"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	err := godotenv.Load()
	connect := os.Getenv("PG_URL")
	db, err := gorm.Open("postgres", connect)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}
	defer db.Close()
	db.AutoMigrate(&dexter.Alert{})
	db.AutoMigrate(&dexter.Webhook{})
}
