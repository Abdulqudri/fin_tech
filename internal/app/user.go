package app

import (
	userhd "github.com/Abdulqudri/fintech/internal/delivery/http/user"
	"github.com/Abdulqudri/fintech/internal/infrastructure/gorm_repo"
	useruc "github.com/Abdulqudri/fintech/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

func (a *App) MountUser(rg *gin.RouterGroup) {
	user_repo := gormrepo.NewRepository(a.DB)
	wallet_repo := gormrepo.NewWalletRepository(a.DB)

	service := useruc.NewService(user_repo, wallet_repo, a.DB)
	handler := userhd.NewHandler(service)
 	userhd.Mount(rg, handler)
}