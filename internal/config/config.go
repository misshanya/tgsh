package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	BotToken    string
	AllowedUser int64
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Info("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		return nil, errors.New("BOT_TOKEN environment variable not set")
	}

	allowedUserStr := os.Getenv("ALLOWED_USER")
	if allowedUserStr == "" {
		return nil, errors.New("ALLOWED_USER environment variable not set")
	}
	allowedUser, err := strconv.Atoi(allowedUserStr)
	if err != nil {
		return nil, errors.New("ALLOWED_USER environment variable is not int64")
	}

	return &Config{
		BotToken:    botToken,
		AllowedUser: int64(allowedUser),
	}, nil
}
