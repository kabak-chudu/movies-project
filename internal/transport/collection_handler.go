package transport

import (
	"github.com/gin-gonic/gin"
	"movies/internal/services"
	"movies/internal/repository"
	"movies/internal/services"
)

type CollectionHandler struct {
	service services.CollectionService
}

func NewCollectionHandler(service services.CollectionService) *CollectionHandler {
	return &CollectionHandler{service: service}
}

func (h *CollectionHandler) RegisterRoutes(r *gin.Engine) {
	collections := r.Collection("/collections")
	{
		collection.Get("", h.GetAll)
		collection.Get("", h.Create)
		collection.Get("/:id", h.GetByID)
		collection.Get("/:id/movies", h.AddMovie)
		collection.Get("/:id/movies/:movie_id", h.RemoveMovie)
	}
}