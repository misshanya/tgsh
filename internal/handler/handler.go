package handler

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/misshanya/tgsh/internal/service"
	"log/slog"
)

type Handler struct {
	service     service.Service
	allowedUser int64
}

func NewHandler(service service.Service, allowedUser int64) *Handler {
	return &Handler{
		service:     service,
		allowedUser: allowedUser,
	}
}

func (h *Handler) StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Hi! This is TGSH. Tool that provides you easier access to your system's shell.",
	})
	if err != nil {
		slog.Error("Error sending message", update.Message.Chat.ID, err)
		return
	}
}

func (h *Handler) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.Chat.ID != h.allowedUser {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "You are not allowed to use this command.",
		})
		if err != nil {
			slog.Error("Error sending message", update.Message.Chat.ID, err)
		}
		return
	}
	out, err := h.service.ExecuteCommand(ctx, update.Message.Text)
	if err != nil {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("[!] Error: %s", err.Error()),
		})
		if err != nil {
			slog.Error("Error sending message", update.Message.Chat.ID, err)
		}
		return
	}
	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   out,
	})
	if err != nil {
		slog.Error("Error sending message", update.Message.Chat.ID, err)
	}
}
