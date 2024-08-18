package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

type Logger struct {
	Log *slog.Logger
}

func LogInit(modeLog string) *Logger {
	var handler slog.Handler
	var logger *slog.Logger
	switch modeLog {

	case "debug":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case "jsonDebug":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case "jsonInfo":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	}

	logger = slog.New(handler)
	slog.SetDefault(logger)

	return &Logger{
		Log: logger,
	}
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func (l *Logger) LogRequest(r *http.Request, requestID string) {
	l.Log.Info("Incoming request",
		slog.String("request_id", requestID),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("query", r.URL.RawQuery),
	)
}

func (l *Logger) LogResponse(r *http.Request, status int, requestID string, responseBody string) {
	level := slog.LevelInfo
	if status >= 400 && status < 500 {
		level = slog.LevelWarn
	} else if status >= 500 {
		level = slog.LevelError
	}

	l.Log.Log(context.Background(), level, "Response",
		slog.String("request_id", requestID),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.Int("status", status),
		slog.String("response", responseBody),
	)
}

func (l *Logger) LogError(ctx context.Context, err error) {
	requestID := ctx.Value("request_id").(string)
	l.Log.Error("An error occurred",
		Err(err),
		slog.String("request_id", requestID),
	)
}
