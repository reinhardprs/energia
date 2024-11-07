package main

import (
	"energia/config"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()

	db, _ := config.ConnectDatabase()
	config.MigrateDB(db)

	e := echo.New()

	e.Start(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Failed loading .env file")
	}
}
