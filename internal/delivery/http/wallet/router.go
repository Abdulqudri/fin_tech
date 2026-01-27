package wallet


import "github.com/gin-gonic/gin"	

func Mount(r *gin.RouterGroup, h *Handler) {

	userGroup := r.Group("/wallets")
	userGroup.GET("/:id", h.GetById)
}	
