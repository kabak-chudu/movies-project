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

	if err := db.AutoMigrate(&models.Movie{}, &models.Genre{}, &models.Review{}); err != nil {
		panic(err)
	}

	movieRepo := repository.NewMovieRepository(db, logger)
	collectionRepo := repository.NewCollectionRepository(db)
	generRepo := repository.NewGenereRepository(db)
	reviewsRepo := repository.NewReviewRepository(db)

	movieService := services.NewMovieService(movieRepo, generRepo, logger)
	collectionService := services.NewCollectionService(collectionRepo, movieRepo)
	generService := services.NewGenereteService(generRepo)
	reviewService := services.NewReviewService(reviewsRepo)

	router := gin.Default()
	transport.RegisterRoutes(router, movieService, collectionService, generService, reviewService, logger)

	port := ":8080"
	logger.Info("server started",
		"addr=", port,
		"env=", env,
	)
	if err := router.Run(port); err != nil {
		panic(err)
	}

}
