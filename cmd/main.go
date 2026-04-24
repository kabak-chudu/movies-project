package main

import (
	"log/slog"
	"movies/internal/config"
	"movies/internal/models"
	"movies/internal/repository"
	"movies/internal/services"
	"movies/internal/transport"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	db := config.SetDatabaseConnection()

	env := os.Getenv("ENV")

	if err := db.AutoMigrate(&models.Movie{}); err != nil {
		panic(err)
	}

	movieRepo := repository.NewMovieRepository(db, logger)

	movieService := services.NewMovieService(movieRepo, logger)

	router := gin.Default()
	transport.RegisterRoutes(router, movieService, logger)

	port := ":8080"
	logger.Info("server started",
		"addr=", port,
		"env=", env,
	)
	if err := router.Run(port); err != nil {
		panic(err)
	}

}
