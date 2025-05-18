package config

import (
	"context"
	"log/slog"
	"os"
)

const (
	TraceIDContextKey = "traceId"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if requestID, ok := ctx.Value(TraceIDContextKey).(string); ok {
		r.AddAttrs(slog.String(TraceIDContextKey, requestID))
	}

	return h.Handler.Handle(ctx, r)
}

func ConfigureLogger(appConfig AppConfig) {
	defaultAttrs := []slog.Attr{
		slog.String("service", appConfig.Name),
		slog.String("environment", appConfig.Environment),
		slog.String("version", appConfig.Version),
	}
	handlerOptions := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}

	baseHandler := slog.NewJSONHandler(os.Stderr, handlerOptions).WithGroup("metadata").WithAttrs(defaultAttrs)
	customHandler := &ContextHandler{baseHandler}
	logger := slog.New(customHandler)

	slog.SetDefault(logger)
}
