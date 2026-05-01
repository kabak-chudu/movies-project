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

	if err := db.AutoMigrate(&models.User{}, &models.Genre{}, &models.Movie{}, &models.Review{}, &models.Collection{}, &models.Watchlist{}); err != nil {
		panic(err)
	}

	movieRepo := repository.NewMovieRepository(db, logger)
	collectionRepo := repository.NewCollectionRepository(db)
	generRepo := repository.NewGenereRepository(db)
	reviewsRepo := repository.NewReviewRepository(db)
	registerRepo := repository.NewRegisterRepository(db)
	loginRepo := repository.NewLoginRepository(db)

	movieService := services.NewMovieService(movieRepo, generRepo, logger)
	collectionService := services.NewCollectionService(collectionRepo, movieRepo)
	generService := services.NewGenereteService(generRepo)
	reviewService := services.NewReviewService(reviewsRepo)
	registerService := services.NewRegisterService(registerRepo)
	loginService := services.NewLoginService(loginRepo)

	router := gin.Default()
	transport.RegisterRoutes(router, movieService, collectionService, generService, reviewService, registerService, loginService, logger)

	port := ":8080"
	logger.Info("server started",
		"addr=", port,
		"env=", env,
	)
	if err := router.Run(port); err != nil {
		panic(err)
	}

}
