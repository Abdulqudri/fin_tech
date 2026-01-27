package wallet

import (
	walletuc "github.com/Abdulqudri/fintech/internal/usecase/wallet"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service *walletuc.WalletService
}

func NewHandler(service *walletuc.WalletService) *Handler {
	return &Handler{service: service}
}

func (h *Handler)GetById(c *gin.Context) {
	ctx := c.Request.Context()
	id := uuid.MustParse(c.Param("id"))
	
	wallet, err := h.service.GetById(ctx, id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Wallet not found"})
		return
	}
	c.JSON(200, wallet)

}