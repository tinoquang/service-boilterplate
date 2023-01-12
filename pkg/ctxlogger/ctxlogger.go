package ctxlogger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxMarker struct{}

type ctxLogger struct {
	logger *zap.Logger
	fields []zapcore.Field
}

var (
	defaultLogger *zap.Logger

	ctxMarkerKey = &ctxMarker{}
)

func SetDefaultLogger(l *zap.Logger) {
	defaultLogger = l
}

// AddFields adds zap fields to the logger.
func AddFields(ctx context.Context, fields ...zapcore.Field) {
	l, ok := ctx.Value(ctxMarkerKey).(*ctxLogger)
	if !ok || l == nil {
		return
	}
	l.fields = append(l.fields, fields...)
}

// Extract takes the call-scoped Logger from the context.
func Extract(ctx context.Context) *zap.Logger {
	l, ok := ctx.Value(ctxMarkerKey).(*ctxLogger)
	if !ok || l == nil {
		return defaultLogger
	}

	return l.logger
}

// ToContext adds the zap.Logger to the context for extraction later.
// Returning the new context that has been created.
func ToContext(ctx context.Context, logger *zap.Logger) context.Context {
	l := &ctxLogger{
		logger: logger,
	}
	return context.WithValue(ctx, ctxMarkerKey, l)
}

// Debug is equivalent to calling Debug on the zap.Logger in the context.
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	Extract(ctx).WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

// Info is equivalent to calling Info on the zap.Logger in the context.
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	Extract(ctx).WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

// Warn is equivalent to calling Warn on the zap.Logger in the context.
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	Extract(ctx).WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

// Error is equivalent to calling Error on the zap.Logger in the context.
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	Extract(ctx).WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}
