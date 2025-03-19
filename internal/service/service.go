package service

import (
	"context"
	"log/slog"
	"os/exec"
	"runtime"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type Service interface {
	ExecuteCommand(ctx context.Context, command string) (string, error)
	executeCommandUnix(ctx context.Context, command string) (string, error)
	executeCommandWindows(ctx context.Context, command string) (string, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) ExecuteCommand(ctx context.Context, command string) (string, error) {
	if runtime.GOOS == "windows" {
		return s.executeCommandWindows(ctx, command)
	}
	return s.executeCommandUnix(ctx, command)
}

func (s *service) executeCommandUnix(ctx context.Context, command string) (string, error) {
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("Error executing command", slog.String("command", command), slog.Any("error", err), slog.String("output", string(out)))
		return "", err
	}
	return string(out), nil
}

func (s *service) executeCommandWindows(ctx context.Context, command string) (string, error) {
	cmd := exec.CommandContext(ctx, "cmd", "/C", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("Error executing command", slog.String("command", command), slog.Any("error", err), slog.String("output", string(out)))
		return "", err
	}

	// From Windows1251 to UTF-8
	decoder := charmap.Windows1251.NewDecoder()
	utf8out, _, _ := transform.String(decoder, string(out))
	return utf8out, nil
}
