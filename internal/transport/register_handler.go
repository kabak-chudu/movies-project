package transport

import (
	"movies/internal/models"
	"movies/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	register services.RegisterService
}

func NewRegisterHandler(
	register services.RegisterService,
) *RegisterHandler {
	return &RegisterHandler{register: register}
}

func (h *RegisterHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/register", func(ctx *gin.Context) {
		var req models.UserCreateRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := h.register.Register(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"created": user})
	})
}
