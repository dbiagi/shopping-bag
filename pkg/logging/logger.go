package logging

import (
    "context"
    "log/slog"
    "os"
    "strings"
)

type Logger interface {
    Log(ctx context.Context, level slog.Level, msg string, args ...any)
    LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr)
    Debug(msg string, args ...any)
    DebugContext(ctx context.Context, msg string, args ...any)
    Info(msg string, args ...any)
    InfoContext(ctx context.Context, msg string, args ...any)
    Warn(msg string, args ...any)
    WarnContext(ctx context.Context, msg string, args ...any)
    Error(msg string, args ...any)
    ErrorContext(ctx context.Context, msg string, args ...any)
}

// NewLogger returns a slog.Logger with the default handler.
func NewLogger(level slog.Level, attrs map[string]string) Logger {
    opts := &slog.HandlerOptions{
        AddSource: true,
        Level:     level,
    }

    var slogAttrs []slog.Attr
    for k, v := range attrs {
        slogAttrs = append(slogAttrs, slog.String(k, v))
    }

    h := slog.NewJSONHandler(os.Stderr, opts).
        WithGroup("context").
        WithAttrs(slogAttrs)

    logger := slog.New(h)
    slog.SetDefault(logger)

    return logger
}

func GetDefaultLogger() Logger {
    return slog.Default()
}

func LogLevelFromString(level string) slog.Level {
    switch strings.ToUpper(level) {
    case "DEBUG":
        return slog.LevelDebug
    case "INFO":
        return slog.LevelInfo
    case "WARN":
        return slog.LevelWarn
    case "ERROR":
        return slog.LevelError
    default:
        return slog.LevelInfo
    }
}
