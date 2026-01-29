package auth

import (
	"net/http"

	authuc "github.com/Abdulqudri/fintech/internal/usecase/auth"
	"github.com/Abdulqudri/fintech/internal/utils/security"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *authuc.AuthService
	jwt     *security.JWTIssuer
}

func NewHandler(serv *authuc.AuthService, jwt *security.JWTIssuer) *Handler {
	return &Handler{
		service: serv,
		jwt: jwt,
	}
}

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var login LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	session, err := h.service.Login(ctx, login.Email, login.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
	}
	accessToken, _ := h.jwt.IssueAccessToken(session.UserID)
	refreshToken, _ := h.service.IssueRefreshToken(ctx, session.UserID)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
