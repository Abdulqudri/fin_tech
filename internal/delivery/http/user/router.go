package user

import "github.com/gin-gonic/gin"	

func Mount(r *gin.RouterGroup, h *Handler) {

	userGroup := r.Group("/users")
	userGroup.POST("", h.Create)
	userGroup.GET("/:email", h.GetByEmail)
}	
