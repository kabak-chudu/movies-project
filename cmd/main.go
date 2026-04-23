package main

import (
	"movies/internal/config"
	"movies/internal/models"
	"movies/internal/repository"
	"movies/internal/services"
	"movies/internal/transport"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetDatabaseConnection()

	if err := db.AutoMigrate(&models.Movie{}); err != nil {
		panic(err)
	}

	movieRepo := repository.NewMovieRepository(db)

	movieService := services.NewMovieService(movieRepo)

	router := gin.Default()
	transport.RegisterRoutes(router, movieService)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
