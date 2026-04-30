package transport

import (
	"movies/internal/models"
	"movies/internal/services"
	"net/http"
	"errors"
	"strconv"
	"github.com/gin-gonic/gin"
)

type CollectionHandler struct {
	collection services.CollectionService
}

func NewCollectionHandler(collection services.CollectionService) *CollectionHandler {
	return &CollectionHandler{collection: collection}
}

func (h *CollectionHandler) RegisterRoutes(router *gin.Engine) {
	collections := router.Group("/collections")
	{
		collections.GET("", h.GetAll)
		collections.POST("", h.Create)
		collections.GET("/:id", h.GetByID)
		collections.POST("/:id/movies", h.AddMovie)
		collections.DELETE("/:id/movies/:movie_id", h.RemoveMovie)
	}
}

func (h *CollectionHandler) Create(c *gin.Context) {
	var req models.CollectionCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection, err := h.collection.CreateCollection(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, collection)
}

func (h *CollectionHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID коллекции"})
		return
	}
	
	collection, err := h.collection.GetCollectionByID(uint(id))
	if err != nil {
		if errors.Is(err, services.ErrCollectionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

	c.JSON(http.StatusOK, collection)
}

func (h *CollectionHandler) GetAll(c *gin.Context) {
	collections, err := h.collection.GetAllCollections()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, collections)
}

func (h *CollectionHandler) AddMovie(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID коллекции"})
		return
	}
	
	var req models.CollectionAddRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection, err := h.collection.AddMovieToCollection(uint(id), req); if err != nil {
		switch {
        	case errors.Is(err, services.ErrCollectionNotFound):
        	    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        	case errors.Is(err, services.ErrMovieNotFound):
        	    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        	default:
        	    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	}
        return
	}

	c.JSON(http.StatusOK, collection)
}

func (h *CollectionHandler) RemoveMovie(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID коллекции"})
		return
	}

	mID, err := strconv.ParseUint(c.Param("movie_id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID фильма"})
        return
    }

	if err := h.collection.RemoveMovieFromCollection(uint(id), uint(mID)); err != nil {
		switch {
        	case errors.Is(err, services.ErrCollectionNotFound):
        	    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        	case errors.Is(err, services.ErrMovieNotFound):
        	    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        	default:
        	    c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
        }
        return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}