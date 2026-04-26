package transport

import (
	"log/slog"
	"movies/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	movie services.MovieService,
	logger *slog.Logger,
) {
	movieHandler := NewMovieHandler(movie, logger)
	movieHandler.RegisterRoutes(router)
}
