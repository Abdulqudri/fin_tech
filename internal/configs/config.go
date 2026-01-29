package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbURL  string
	Port      string
	JwtSecret string
}

func Load() (*Config, error) {
	err :=godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading configuration from environment variables")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is required")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	secret := os.Getenv("JWT_SECRETE")
	
	return &Config{
		DbURL: dbUrl,
		Port:  port,
		JwtSecret: secret,
	}, nil
}
