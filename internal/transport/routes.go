package transport

import (
	"movies/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	movie services.MovieService,
) {
	movieHandler := NewMovieHandler(movie)
	movieHandler.RegisterRoutes(router)
}
