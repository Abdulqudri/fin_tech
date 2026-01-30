package configs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DbURL  string
	Port      string
	JwtSecret string
	JwtExpiry time.Duration
	RefreshExpiry time.Duration
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
	jwtExpiry := os.Getenv("JWT_EXPIRY")
	refreshExpiry := os.Getenv("REFRESH_EXPIRY")
	log.Println(jwtExpiry)
	log.Println(refreshExpiry)

	jwtDuration, err := time.ParseDuration(jwtExpiry)
	if err != nil {
		log.Fatal("Invalid JWT_EXPIRY value")
	}
	refreshDuration, err := time.ParseDuration(refreshExpiry)
	if err != nil {
		log.Fatal("Invalid REFRESH_EXPIRY value")
	}

	return &Config{
		DbURL: dbUrl,
		Port:  port,
		JwtSecret: secret,
		JwtExpiry: jwtDuration,
		RefreshExpiry: refreshDuration,
	}, nil
}

