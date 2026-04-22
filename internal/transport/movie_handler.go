package transport

import (
	"errors"
	"log/slog"
	"movies/internal/models"
	"movies/internal/repository"
	"movies/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieHandler struct {
	movie  services.MovieService
	logger *slog.Logger
}

func NewMovieHandler(
	movie services.MovieService,
	logger *slog.Logger,
) *MovieHandler {
	return &MovieHandler{movie: movie, logger: logger}
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
	h.logger.Info("incoming request",
		"layer", "handler",
		"level", "Info",
		"method", ctx.Request.Method,
		"path", ctx.Request.URL.Path,
	)
	var req models.CreateMovieRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("validate req fail",
			"layer", "handler",
			"level", "Warn",
			"reason", err.Error(),
			"title", req.Title,
			"year", req.Year,
			"country", req.Country,
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.movie.CreateMovie(&req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.logger.Error("create movie error",
				"layer", "handler",
				"level", "Error",
				"error", err.Error(),
			)
			ctx.JSON(http.StatusNotFound, gin.H{"error": "такого жанра по айди не существует"})
			return
		}
		h.logger.Error("create movie error",
			"layer", "handler",
			"error", err.Error(),
			"level", "Error",
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("succesful created",
		"layer", "handler",
		"status_code", http.StatusCreated,
		"movie_id", movie.ID,
	)
	ctx.JSON(http.StatusCreated, gin.H{"created": movie})
}

func (h *MovieHandler) GetAllMovies(ctx *gin.Context) {
	h.logger.Info("incoming request",
		"layer", "handler",
		"level", "Info",
		"method", ctx.Request.Method,
		"path", ctx.Request.URL.Path,
	)
	var filter repository.MovieFilter

	genreQuery := ctx.Query("genre_id")
	if genreQuery != "" {
		genreID, err := strconv.Atoi(genreQuery)
		if err != nil {
			h.logger.Warn("validate id",
				"layer", "handler",
				"level", "Warn",
				"reason", err.Error(),
				"genre_id", genreID,
			)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			h.logger.Error("fail get movies:",
				"layer", "handler",
				"error", err.Error(),
				"level", "Error",
			)
			return
		}
		genreId := uint(genreID)
		filter.GenreID = &genreId
	}
	yearQuery := ctx.Query("year")
	if yearQuery != "" {
		yearr, err := strconv.Atoi(yearQuery)
		if err != nil {
			h.logger.Warn("convert query fail",
				"layer", "handler",
				"reason", err.Error(),
				"year", yearr,
			)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			h.logger.Error("Error:",
				"layer", "handler",
				"error", err.Error(),
			)
			return
		}
		filter.Year = &yearr
	}

	movies, err := h.movie.GetAllByFilter(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logger.Error("movies get fail",
			"layer", "handler",
			"error", err.Error(),
		)
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"movies": movies})
	h.logger.Info("succes get movies",
		"layer", "handler",
		"status", http.StatusOK,
	)
}

func (h *MovieHandler) GetMovieByID(ctx *gin.Context) {
	h.logger.Info("incoming request",
		"layer", "handler",
		"method", ctx.Request.Method,
		"path", ctx.Request.URL.Path,
		"raw_id", ctx.Param("id"),
	)
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Warn("invalid movie id format",
			"layer", "handler",
			"reason", err.Error(),
			"input_value", idStr,
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.movie.GetMovieByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.logger.Warn("movie not found", "id", id, "layer", "handler")
			ctx.JSON(http.StatusNotFound, gin.H{"error": services.ErrMovieNotFound.Error()})
			return
		}
		h.logger.Error("failed to get movie",
			"layer", "handler",
			"error", err.Error(),
			"id", id,
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	h.logger.Info("request processed",
		"layer", "handler",
		"status", http.StatusOK,
		"movie_id", movie.ID,
	)
	ctx.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) PatchMovieByID(ctx *gin.Context) {
	h.logger.Info("incoming request",
		"layer", "handler",
		"method", ctx.Request.Method,
		"path", ctx.Request.URL.Path,
		"raw_id", ctx.Param("id"),
	)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logger.Warn("Warn:",
			"layer", "handler",
			"reason", err.Error(),
			"input_value", ctx.Param("id"),
		)
		return
	}
	var req *models.UpdateMovieRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Warn:",
			"layer", "handler",
			"reason", err.Error(),
			"Country", req.Country,
			"year", req.Year,
			"title", req.Title,
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.movie.UpdatePATCHMovie(uint(id), req)
	if err != nil {
		if err.Error() == "такого фильма по айди не существует" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			h.logger.Warn("movie not found", "id", id, "layer", "handler")
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logger.Error("failed to update movie",
			"layer", "handler",
			"error", err.Error(),
			"id", id,
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"edited": movie})
	h.logger.Info("request processed:",
		"status", http.StatusOK,
		"layer", "handler",
		"movie_id", movie.ID,
	)

}

func (h *MovieHandler) DeleteMovieByID(ctx *gin.Context) {
	h.logger.Info("incoming request",
		"layer", "handler",
		"method", ctx.Request.Method,
		"path", ctx.Request.URL.Path,
		"raw_id", ctx.Param("id"),
	)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		h.logger.Warn("invalid movie id format",
			"reason", err.Error(),
			"layer", "handler",
			"input_value", ctx.Param("id"),
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.movie.DeleteMovie(uint(id)); err != nil {

		h.logger.Error("failed to delete movie",
			"error", err.Error(),
			"layer", "handler",
			"id", id,
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
	h.logger.Info("succesful deleted",
		"layer", "handler",
		"status", http.StatusOK,
		"message", "deleted",
	)
}
