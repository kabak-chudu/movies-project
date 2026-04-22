package transport

import (
	"movies/internal/models"
	"movies/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GenerHandler struct {
	gener services.GenereService
}

func NewGenerHandler(
	gener services.GenereService,
) *GenerHandler {
	return &GenerHandler{gener: gener}
}

func (h *GenerHandler) RigisterRoutes(router *gin.Engine) {
	gener := router.Group("/generes")
	{
		gener.POST("/", h.CreateGenere)
		gener.GET("/", h.GetAllGeneres)
		gener.GET("/:id", h.GetGenerByID)
		gener.PATCH("/:id", h.UpdatePATCHGener)
		gener.DELETE("/:id", h.DeleteGener)
	}
}

func (h *GenerHandler) CreateGenere(ctx *gin.Context) {
	var req models.CreateGenreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gener, err := h.gener.CreateGenere(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gener)
}

func (h *GenerHandler) GetGenerByID(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	genre, err := h.gener.GetGenerByID(uint(idParam))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, genre)
}

func (h *GenerHandler) GetAllGeneres(ctx *gin.Context) {
	generes, err := h.gener.GetAllGeneres()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, generes)
}

func (h *GenerHandler) UpdatePATCHGener(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req *models.CreateGenreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gener, err := h.gener.UpdatePATCHGener(uint(idParam), &models.UpdateGenreRequest{Name: req.Name})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gener)
}

func (h *GenerHandler) DeleteGener(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.gener.DeleteGener(uint(idParam)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"messege": "Успешно удалена!"})

}
