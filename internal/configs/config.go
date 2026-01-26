package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbURL  string
	Port      string
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
	
	return &Config{
		DbURL: dbUrl,
		Port:  port,
	}, nil
}
