package main

import (
	"movies/internal/config"
	"movies/internal/models"
)

func main() {
	db := config.SetDatabaseConnection()

	if err := db.AutoMigrate(&models.Genre{}, &models.Movie{}); err != nil {
		panic(err)
	}
}
