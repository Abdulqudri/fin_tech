package app

import (
	wallethd"github.com/Abdulqudri/fintech/internal/delivery/http/wallet"
	gormrepo "github.com/Abdulqudri/fintech/internal/infrastructure/gorm_repo"
	walletuc "github.com/Abdulqudri/fintech/internal/usecase/wallet"
	"github.com/gin-gonic/gin"
)

func (a *App) MountWallet(rg *gin.RouterGroup) {
	user_repo := gormrepo.NewRepository(a.DB)
	wallet_repo := gormrepo.NewWalletRepository(a.DB)

	service := walletuc.NewService(wallet_repo, user_repo)
	handler := wallethd.NewHandler(service)
	wallethd.Mount(rg, handler)
}