package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger defines the logging interface for the application
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)

	// Context-aware logging
	DebugContext(ctx context.Context, msg string, fields ...Field)
	InfoContext(ctx context.Context, msg string, fields ...Field)
	WarnContext(ctx context.Context, msg string, fields ...Field)
	ErrorContext(ctx context.Context, msg string, fields ...Field)

	// With creates a child logger with additional fields
	With(fields ...Field) Logger

	// Sync flushes any buffered log entries
	Sync() error
}

// Field represents a structured logging field
type Field = zap.Field

// ZapLogger wraps zap.Logger to implement our Logger interface
type ZapLogger struct {
	logger *zap.Logger
}

// Config holds logger configuration
type Config struct {
	Level       string `json:"level"`       // debug, info, warn, error
	Environment string `json:"environment"` // development, production
	OutputPath  string `json:"output_path"` // stdout, stderr, or file path
}

// New creates a new logger instance
func New(cfg Config) (Logger, error) {
	level, err := parseLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}

	var zapConfig zap.Config

	switch cfg.Environment {
	case "development":
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case "production":
		zapConfig = zap.NewProductionConfig()
	default:
		zapConfig = zap.NewDevelopmentConfig()
	}

	zapConfig.Level = zap.NewAtomicLevelAt(level)

	if cfg.OutputPath != "" {
		zapConfig.OutputPaths = []string{cfg.OutputPath}
	}

	logger, err := zapConfig.Build(
		zap.AddCallerSkip(1), // Skip the wrapper function
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}

	return &ZapLogger{logger: logger}, nil
}

// NewDefault creates a logger with sensible defaults
func NewDefault() Logger {
	cfg := Config{
		Level:       "info",
		Environment: "development",
		OutputPath:  "stdout",
	}

	logger, err := New(cfg)
	if err != nil {
		// Fallback to a basic logger if configuration fails
		zapLogger, _ := zap.NewDevelopment()
		return &ZapLogger{logger: zapLogger}
	}

	return logger
}

// Debug logs a debug message
func (l *ZapLogger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

// Info logs an info message
func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

// Warn logs a warning message
func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

// Error logs an error message
func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

// Fatal logs a fatal message and exits
func (l *ZapLogger) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

// DebugContext logs a debug message with context
func (l *ZapLogger) DebugContext(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, contextFields(ctx)...)
	l.logger.Debug(msg, fields...)
}

// InfoContext logs an info message with context
func (l *ZapLogger) InfoContext(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, contextFields(ctx)...)
	l.logger.Info(msg, fields...)
}

// WarnContext logs a warning message with context
func (l *ZapLogger) WarnContext(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, contextFields(ctx)...)
	l.logger.Warn(msg, fields...)
}

// ErrorContext logs an error message with context
func (l *ZapLogger) ErrorContext(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, contextFields(ctx)...)
	l.logger.Error(msg, fields...)
}

// With creates a child logger with additional fields
func (l *ZapLogger) With(fields ...Field) Logger {
	return &ZapLogger{logger: l.logger.With(fields...)}
}

// Sync flushes any buffered log entries
func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}

// Helper functions

func parseLevel(level string) (zapcore.Level, error) {
	switch level {
	case "debug":
		return zapcore.DebugLevel, nil
	case "info":
		return zapcore.InfoLevel, nil
	case "warn":
		return zapcore.WarnLevel, nil
	case "error":
		return zapcore.ErrorLevel, nil
	default:
		return zapcore.InfoLevel, fmt.Errorf("unknown level: %s", level)
	}
}

// contextFields extracts logging fields from context
func contextFields(ctx context.Context) []Field {
	var fields []Field

	// Extract common context values for logging
	if userID := ctx.Value("user_id"); userID != nil {
		fields = append(fields, zap.Any("user_id", userID))
	}

	if chatID := ctx.Value("chat_id"); chatID != nil {
		fields = append(fields, zap.Any("chat_id", chatID))
	}

	if requestID := ctx.Value("request_id"); requestID != nil {
		fields = append(fields, zap.String("request_id", requestID.(string)))
	}

	return fields
}

// Convenience field constructors (re-export zap functions for ease of use)
var (
	String   = zap.String
	Int      = zap.Int
	Int64    = zap.Int64
	Float64  = zap.Float64
	Bool     = zap.Bool
	Duration = zap.Duration
	Time     = zap.Time
	Error    = zap.Error
	Any      = zap.Any
)
