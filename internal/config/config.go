package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	APIKey   string
	BotToken string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Cannot load .env file: %s", err.Error())
	}

	return &Config{
		APIKey:   os.Getenv("API_KEY"),
		BotToken: os.Getenv("BOT_TOKEN"),
	}
}
