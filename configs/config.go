package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB Dbconfig
}

type Dbconfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Print("Error loading .env file")
	}

	return &Config{
		DB: Dbconfig{
			Dsn: os.Getenv("DSN"),
		},
	}

}
