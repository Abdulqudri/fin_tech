package wallet

import (
	middleware "github.com/Abdulqudri/fintech/internal/delivery/middlewares"
	"github.com/Abdulqudri/fintech/internal/utils/security"
	"github.com/gin-gonic/gin"
)	

func Mount(r *gin.RouterGroup, h *Handler) {

	api := r.Group("/wallets")
	api.Use(middleware.JWTAuth(&security.JWTIssuer{}))
	api.GET("/:id", h.GetById)
	api.GET("/user/:id", h.GetByUserId)
}	
