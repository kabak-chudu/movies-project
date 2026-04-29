package transport

import (
	"movies/internal/models"
	"movies/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	review services.ReviewService
}

func NewReviewHandler(
	review services.ReviewService,
) *ReviewHandler {
	return &ReviewHandler{review: review}
}

func (h *ReviewHandler) RigisterRoutes(router *gin.Engine) {
	reviews := router.Group("/reviews")
	{
		reviews.POST("/", h.CreateReview)
		reviews.GET("/", h.GetReviewAll)
		reviews.GET("/:id", h.GetReviewByID)
		reviews.PATCH("/:id", h.UpdatePATCHReview)
		reviews.DELETE("/:id", h.DeleteReview)
	}
}

func (h *ReviewHandler) CreateReview(ctx *gin.Context) {
	var req models.CreateReviewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	review, err := h.review.CreateReview(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, review)
}

func (h *ReviewHandler) GetReviewAll(ctx *gin.Context) {
	review, err := h.review.GetReviewAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) GetReviewByID(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	review, err := h.review.GetReviewByID(uint(idParam))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) DeleteReview(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.review.DeleteReview(uint(idParam)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Запись успешно удалена!"})
}

func (h *ReviewHandler) UpdatePATCHReview(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req *models.UpdateReviewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	review, err := h.review.UpdateReview(uint(idParam), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, review)
}
