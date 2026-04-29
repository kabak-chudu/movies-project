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
	review services.ReviewService,
	genere services.GenereService,
) {
	movieHandler := NewMovieHandler(movie, logger)
	movieHandler.RegisterRoutes(router)
	reviewHandler := NewReviewHandler(review)
	reviewHandler.RigisterRoutes(router)
	genereHandler := NewGenerHandler(genere)
	genereHandler.RigisterRoutes(router)
}
