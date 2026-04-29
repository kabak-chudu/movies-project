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
	logger *slog.Logger,
) {
	movieHandler := NewMovieHandler(movie, logger)
	collectionHandler := NewCollectionHandler(collection)

	movieHandler.RegisterRoutes(router)
	collectionHandler.RegisterRoutes(router)
}