package app

import (

	authhd "github.com/Abdulqudri/fintech/internal/delivery/http/auth"
	gormrepo "github.com/Abdulqudri/fintech/internal/infrastructure/gorm_repo"
	authuc "github.com/Abdulqudri/fintech/internal/usecase/auth"
	"github.com/Abdulqudri/fintech/internal/utils/security"
	"github.com/gin-gonic/gin"
)

func (a *App) MountAuth(rg *gin.RouterGroup) {
	auth_repo := gormrepo.NewAuthRepository(a.DB)
	service := authuc.NewService(auth_repo, a.Config.RefreshExpiry)
	jwt := security.NewJWTIssuer(a.Config.JwtSecret, a.Config.JwtExpiry)
	handler := authhd.NewHandler(service, jwt)
	authhd.Mount(rg, handler)
}