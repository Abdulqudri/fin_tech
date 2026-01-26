package app

import (
	"github.com/Abdulqudri/fintech/internal/configs"
	"github.com/Abdulqudri/fintech/internal/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	config 		*configs.Config
	db          *gorm.DB
}

func NewApp() *App {
	config, _ := configs.Load()
	db := db.New(config.DbURL)
	return &App{config: config, db: db}
}

func (a *App) Mount() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
	return r
}