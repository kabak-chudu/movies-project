package transport

import (
	"log/slog"
	"movies/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	movie services.MovieService,
	collection services.CollectionService,
	user services.UserService,
	logger *slog.Logger,
) {
	movieHandler := NewMovieHandler(movie, logger)
	collectionHandler := NewCollectionHandler(collection)
	userHandler := NewUserHandler(user)

	movieHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	collectionHandler.RegisterRoutes(router)
}