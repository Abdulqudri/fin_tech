package auth

import "github.com/gin-gonic/gin"

func Mount(r *gin.RouterGroup, h *Handler) {
	group := r.Group("/auth")
	group.POST("/login", h.Login)
}