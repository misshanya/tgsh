package service

import (
	"context"
	"log/slog"
	"os/exec"
)

type Service interface {
	ExecuteCommand(ctx context.Context, command string) (string, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) ExecuteCommand(ctx context.Context, command string) (string, error) {
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("Error executing command", slog.String("command", command), slog.Any("error", err), slog.String("output", string(out)))
		return "", err
	}
	return string(out), nil
}
