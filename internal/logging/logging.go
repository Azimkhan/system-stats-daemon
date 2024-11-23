package logging

import (
	"errors"
	"log/slog"
	"os"

	"github.com/Azimkhan/system-stats-daemon/internal/config"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Warn(msg string, args ...any)
	With(args ...any) Logger
}

type SlogWrapper struct {
	logger *slog.Logger
}

func (s *SlogWrapper) Debug(msg string, args ...any) {
	s.logger.Debug(msg, args...)
}

func (s *SlogWrapper) Info(msg string, args ...any) {
	s.logger.Info(msg, args...)
}

func (s *SlogWrapper) Error(msg string, args ...any) {
	s.logger.Error(msg, args...)
}

func (s *SlogWrapper) Warn(msg string, args ...any) {
	s.logger.Warn(msg, args...)
}

func (s *SlogWrapper) With(args ...any) Logger {
	return &SlogWrapper{logger: s.logger.With(args...)}
}

var (
	ErrUnknownLogLevel  = errors.New("unknown log level")
	ErrUnknownLogFormat = errors.New("unknown log format")
)

func NewLogger(conf *config.LoggingConfig) (Logger, error) {
	// determine level
	var level slog.Level
	switch conf.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		return nil, ErrUnknownLogLevel
	}
	slog.SetLogLoggerLevel(level)

	// determine handler
	var handler slog.Handler
	switch conf.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})

	case "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	default:
		return nil, ErrUnknownLogFormat
	}
	logger := slog.New(handler)
	return &SlogWrapper{logger: logger}, nil
}
