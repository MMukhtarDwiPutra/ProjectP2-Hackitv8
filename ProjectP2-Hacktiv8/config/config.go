package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	XenditAPIKey string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
	return &Config{
		XenditAPIKey: os.Getenv("3RD_PARTY_XENDIT_API"),
	}
}
