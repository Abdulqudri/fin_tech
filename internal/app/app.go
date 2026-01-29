package app

import (
	"github.com/Abdulqudri/fintech/internal/configs"
	"github.com/Abdulqudri/fintech/internal/infrastructure/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Config 		*configs.Config
	DB          *gorm.DB
}

func NewApp() *App {
	config, _ := configs.Load()
	db := db.New(config.DbURL)
	return &App{Config: config, DB: db}
}

func (a *App) BuildHTTPServer() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	a.MountUser(api)
	a.MountWallet(api)
	a.MountAuth(api)
	return r
}