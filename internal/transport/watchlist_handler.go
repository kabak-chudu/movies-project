package transport

import (
	"errors"
	"movies/internal/models"
	"movies/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WatchListHandler struct {
	watchlist services.WatchlistService
}

func NewWatchListHandler(watchlist services.WatchlistService) *WatchListHandler {
	return &WatchListHandler{watchlist: watchlist}
}

func (h *WatchListHandler) RegisterRoutes(router *gin.Engine) {
	collections := router.Group("/watchlist")
	{
		collections.POST("", h.Create)
		collections.POST("/:id/movie", h.AddMovie)
		collections.DELETE("/:id", h.Remove)
	}
	router.GET("/users/:id/watchlist", h.GetByUserID)
}

func (h *WatchListHandler) Create(ctx *gin.Context) {
	var req models.CreateWatchlistRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	watchlist, err := h.watchlist.CreateWatchList(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": "usera по такому айди не найдено"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, watchlist)
}

func (h *WatchListHandler) GetByUserID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	watchlist, err := h.watchlist.GetWatchListByUserID(uint(id))
	if err != nil {
		if errors.Is(err, services.ErrCollectionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, watchlist)
}

func (h *WatchListHandler) AddMovie(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.WatchlistAddRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	watchlist, err := h.watchlist.AddMovieToWatchList(uint(id), req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, watchlist)
}

func (h *WatchListHandler) Remove(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID коллекции"})
		return
	}

	if err := h.watchlist.RemoveWatchlist(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
