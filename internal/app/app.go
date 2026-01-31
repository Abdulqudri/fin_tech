package app

import (
	"github.com/Abdulqudri/fintech/internal/configs"
	"github.com/Abdulqudri/fintech/internal/infrastructure/db"
	"github.com/Abdulqudri/fintech/internal/infrastructure/models"
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
	app := &App{Config: config, DB: db}
	app.MigrateDB()
	return app
}

func (a *App) BuildHTTPServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	api := r.Group("/api/v1")
	a.MountUser(api)
	a.MountWallet(api)
	a.MountAuth(api)
	return r
}

func (a *App) MigrateDB() {
	err := a.DB.AutoMigrate(
		&models.User{},
		&models.Wallet{},
		&models.RefreshToken{},
		&models.Transaction{},
		&models.LedgerEntry{},
		&models.Transfer{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
}
