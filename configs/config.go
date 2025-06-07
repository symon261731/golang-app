package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Token string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Print("Error loading .env file")
	}

	return &Config{
		DB: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Token: os.Getenv("TOKEN"),
		},
	}

}
