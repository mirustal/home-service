package logger

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var Logger *slog.Logger


func LogInit(modeLog string) {
	var handler slog.Handler

	switch modeLog {
	case "debug":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case "jsonDebug":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case "jsonInfo":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	default:
		log.Fatal("no init modeLog: ", modeLog)
	}

	Logger = slog.New(handler)
	slog.SetDefault(Logger)
}


func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}


func LogRequest(r *http.Request, requestID string) {
	Logger.Info("Incoming request",
		slog.String("request_id", requestID),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.String("query", r.URL.RawQuery),
	)
}


func LogResponse(r *http.Request, status int, requestID string, responseBody string) {
	level := slog.LevelInfo
	if status >= 400 && status < 500 {
		level = slog.LevelWarn
	} else if status >= 500 {
		level = slog.LevelError
	}

	Logger.Log(context.Background(), level, "Response",
		slog.String("request_id", requestID),
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.Int("status", status),
		slog.String("response", responseBody),
	)
}


func LogError(ctx context.Context, err error) {
	requestID := ctx.Value("request_id").(string)
	Logger.Error("An error occurred",
		Err(err),
		slog.String("request_id", requestID),
	)
}
