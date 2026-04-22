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

type RegisterHandler struct {
	register services.RegisterService
	login    services.LoginService
}

func NewRegisterHandler(
	register services.RegisterService,
	login services.LoginService,
) *RegisterHandler {
	return &RegisterHandler{register: register, login: login}
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

	router.POST("/login", func(ctx *gin.Context) {
		var req models.Login

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := h.login.Login(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"успешно вошли в профиль": user})
	})

	router.GET("/users/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := h.login.GetByID(uint(id))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)
	})
}
