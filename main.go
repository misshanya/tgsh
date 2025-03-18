package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/misshanya/tgsh/internal/config"
	"github.com/misshanya/tgsh/internal/handler"
	"github.com/misshanya/tgsh/internal/service"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("Error loading config")
		os.Exit(1)
	}

	// Init services
	botService := service.NewService()

	// Init handlers
	botHandler := handler.NewHandler(botService, cfg.AllowedUser)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(botHandler.DefaultHandler),
	}

	b, err := bot.New(cfg.BotToken, opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, botHandler.StartHandler)

	slog.Info("Bot started")
	b.Start(ctx)
}
