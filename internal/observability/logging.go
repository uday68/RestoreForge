package observability

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(level string) *Logger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})
	return &Logger{
		Logger: slog.New(handler),
	}
}

func (l *Logger) WithTenant(tenantID string) *Logger {
	return &Logger{
		Logger: l.Logger.With(slog.String("tenant_id", tenantID)),
	}
}

func (l *Logger) WithJob(jobID string) *Logger {
	return &Logger{
		Logger: l.Logger.With(slog.String("job_id", jobID)),
	}
}

