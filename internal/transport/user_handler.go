package transport

import (
	"movies/internal/models"
	"movies/internal/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	user services.UserService
}

func NewUserHandler(user services.UserService) *UserHandler {
	return &UserHandler{user: user}
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.POST("", h.Create)
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req models.UserCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.user.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}