package transport

import (
	"errors"
	"movies/internal/models"
	"movies/internal/repository"
	"movies/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieHandler struct {
	movie services.MovieService
}

func NewMovieHandler(
	movie services.MovieService,
) *MovieHandler {
	return &MovieHandler{movie: movie}
}

func (h *MovieHandler) RegisterRoutes(router *gin.Engine) {
	movies := router.Group("/movies")
	{
		movies.POST("/", h.PostMovie)
		movies.GET("/", h.GetAllMovies)
		movies.GET("/:id", h.GetMovieByID)
		movies.PATCH("/:id", h.PatchMovieByID)
		movies.DELETE("/:id", h.DeleteMovieByID)
	}
}

func (h *MovieHandler) PostMovie(ctx *gin.Context) {
	var req models.CreateMovieRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.movie.CreateMovie(&req)
	if err != nil {
		if errors.Is(err, errors.New("такого жанра по айди не существует")) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"created": movie})
}

func (h *MovieHandler) GetAllMovies(ctx *gin.Context) {
	var filter repository.MovieFilter

	genreQuery := ctx.Query("genre_id")
	if genreQuery != "" {
		genreID, err := strconv.Atoi(genreQuery)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		genreId := uint(genreID)
		filter.GenreID = &genreId
	}
	yearQuery := ctx.Query("year")
	if yearQuery != "" {
		yearr, err := strconv.Atoi(yearQuery)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		filter.Year = &yearr
	}

	movies, err := h.movie.GetAllByFilter(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"movies": movies})
}

func (h *MovieHandler) GetMovieByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.movie.GetMovieByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": services.ErrMovieNotFound.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) PatchMovieByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req *models.UpdateMovieRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.movie.UpdatePATCHMovie(uint(id), req)
	if err != nil {
		if err.Error() == "такого фильма по айди не существует" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"edited": movie})

}

func (h *MovieHandler) DeleteMovieByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.movie.DeleteMovie(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
