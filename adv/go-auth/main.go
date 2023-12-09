package main

import (
	"fmt"
	"os"

	"dev.pradeep/packages/models"
)

func main() {
	fmt.Print("im working")
	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	fmt.Print(config)
}
