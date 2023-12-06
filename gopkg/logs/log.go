package logs

import (
	"context"
	"sync"
)

var (
	logger Logger
	mu     sync.Mutex
)

// Setup 修改 logger 默认变量
// 无法保证并发安全，尽量仅在初始化的时候使用
func Setup(l Logger) {
	mu.Lock()
	defer mu.Unlock()
	logger = l
}

type Logger interface {
	Debug(ctx context.Context, format string, args ...any)
	Info(ctx context.Context, format string, args ...any)
	Warn(ctx context.Context, format string, args ...any)
	Error(ctx context.Context, format string, args ...any)
}

func Debug(ctx context.Context, format string, args ...any) {
	logger.Debug(ctx, format, args...)
}

func Info(ctx context.Context, format string, args ...any) {
	logger.Info(ctx, format, args...)
}

func Warn(ctx context.Context, format string, args ...any) {
	logger.Warn(ctx, format, args...)
}

func Error(ctx context.Context, format string, args ...any) {
	logger.Error(ctx, format, args...)
}
