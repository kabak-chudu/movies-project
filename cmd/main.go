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
	// genreRepo := repository.NewGenreRepository(db)

	movieService := services.NewMovieService(movieRepo)

	router := gin.Default()
	transport.RegisterRoutes(router, movieService)
	router.Run(":8080")
}
