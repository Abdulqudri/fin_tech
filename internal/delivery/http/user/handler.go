package user

import (
	"github.com/Abdulqudri/fintech/internal/domain/user"
	useruc "github.com/Abdulqudri/fintech/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *useruc.UserService
}

func NewHandler(service *useruc.UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var u CreateUserRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	user := user.User{
		FullName: u.FullName,
		Email: u.Email,
	}
	password := u.Password
	if err := h.service.CreateUser(ctx, &user, password); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, user)
}
func (h *Handler) GetByEmail(c *gin.Context) {
	ctx := c.Request.Context()
	email := c.Param("email")
	user, err := h.service.GetByEmail(ctx, email)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}
