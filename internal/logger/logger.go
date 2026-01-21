package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Debug logs a debug message using tflog. Prefer passing a request/context from resources/providers.
func Debug(ctx context.Context, msg string, args ...any) {
	tflog.Debug(ctx, fmt.Sprintf(msg, args...))
}

// Info logs an info message using tflog.
func Info(ctx context.Context, msg string, args ...any) {
	tflog.Info(ctx, fmt.Sprintf(msg, args...))
}

// Warn logs a warning using tflog.
func Warn(ctx context.Context, msg string, args ...any) {
	tflog.Warn(ctx, fmt.Sprintf(msg, args...))
}

// Error logs an error using tflog.
func Error(ctx context.Context, msg string, args ...any) {
	tflog.Error(ctx, fmt.Sprintf(msg, args...))
}

var logInstance *slog.Logger

// Backwards-compatible logging functions that do not require a context.

// Using fmt.Sprintf as the formatting is preferable to slog's built-in formatting.
func DebugBG(msg string, args ...any) {
	logInstance.Debug(fmt.Sprintf(msg, args...))
}

func InfoBG(msg string, args ...any) {
	logInstance.Info(fmt.Sprintf(msg, args...))
}

func WarnBG(msg string, args ...any) {
	logInstance.Warn(fmt.Sprintf(msg, args...))
}

func ErrorBG(msg string, args ...any) {
	logInstance.Error(fmt.Sprintf(msg, args...))
}

func setupLogger() *slog.Logger {
	level := os.Getenv("TF_LOG")
	var slogLevel slog.Level
	switch level {
	case "TRACE":
		slogLevel = slog.LevelDebug
	case "DEBUG":
		slogLevel = slog.LevelDebug
	case "INFO":
		slogLevel = slog.LevelInfo
	case "WARN":
		slogLevel = slog.LevelWarn
	case "ERROR":
		slogLevel = slog.LevelError
	default:
		slogLevel = slog.LevelInfo
	}
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slogLevel})

	return slog.New(handler)
}

func init() {
	logInstance = setupLogger()
}
