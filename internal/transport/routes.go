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
	genere services.GenereService,
	review services.ReviewService,
	register services.RegisterService,
	login services.LoginService,
	logger *slog.Logger,
) {
	movieHandler := NewMovieHandler(movie, logger)
	movieHandler.RegisterRoutes(router)

	collectionHandler := NewCollectionHandler(collection)
	collectionHandler.RegisterRoutes(router)

	reviewHandler := NewReviewHandler(review)
	reviewHandler.RigisterRoutes(router)

	genereHandler := NewGenerHandler(genere)
	genereHandler.RigisterRoutes(router)

	registerHandler := NewRegisterHandler(register, login)
	registerHandler.RegisterRoutes(router)
}
