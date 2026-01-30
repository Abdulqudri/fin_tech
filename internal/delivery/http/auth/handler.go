package auth

import (
	"net/http"

	authuc "github.com/Abdulqudri/fintech/internal/usecase/auth"
	"github.com/Abdulqudri/fintech/internal/utils/security"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service *authuc.AuthService
	jwt     *security.JWTIssuer
}

func NewHandler(serv *authuc.AuthService, jwt *security.JWTIssuer) *Handler {
	return &Handler{
		service: serv,
		jwt:     jwt,
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
	accessToken, err := h.jwt.IssueAccessToken(session.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token issuance failed"})
		return
	}
	refreshToken, err := h.service.IssueRefreshToken(ctx, session.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token issuance failed"})
		return
	}
	c.SetCookie(
		"refresh_token",
		refreshToken.ID.String(),
		int(refreshToken.ExpiresAt.Unix()),
		"/",
		"",
		true,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})

}

func (h *Handler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})
		return
	}
	refreshTokenID, err := uuid.Parse(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid refresh token"})
		return
	}
	userId, newRefreshToken, err := h.service.Refresh(ctx, refreshTokenID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}
	accessToken, _ := h.jwt.IssueAccessToken(uuid.MustParse(userId))
	c.SetCookie(
		"refresh_token",
		newRefreshToken.ID.String(),
		int(newRefreshToken.ExpiresAt.Unix()),
		"/",
		"",
		true,
		true,
	)
	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})

}

func (h *Handler) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing refresh token"})
		return
	}
	refreshTokenID, err := uuid.Parse(cookie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid refresh token"})
		return
	}
	_ = h.service.RevokeRefreshToken(ctx, refreshTokenID)
	
	c.SetCookie(
		"refresh_token",
		"",
		0,
		"/",
		"",
		true,
		true,
	)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}