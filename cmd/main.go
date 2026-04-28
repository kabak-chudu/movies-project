package main

import (
	"log/slog"
	"movies/internal/config"
	"movies/internal/models"
	"movies/internal/repository"
	"movies/internal/services"
	"movies/internal/transport"
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	db := config.SetDatabaseConnection()

	env := os.Getenv("ENV")

	if err := db.AutoMigrate(&models.User{}, &models.Movie{}, &models.Collection{}); err != nil {
		panic(err)
	}

	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
    	db.Create(&models.User{Username: "test_user"})
    	fmt.Println("тестовый пользователь создан (ID: 1)")
	}

	movieRepo := repository.NewMovieRepository(db, logger)
	collectionRepo := repository.NewCollectionRepository(db)

	movieService := services.NewMovieService(movieRepo, logger)
	collectionService := services.NewCollectionService(collectionRepo, movieRepo)

	router := gin.Default()
	transport.RegisterRoutes(router, movieService, collectionService, logger)

	port := ":8080"
	logger.Info("server started",
		"addr=", port,
		"env=", env,
	)
	if err := router.Run(port); err != nil {
		panic(err)
	}

}
