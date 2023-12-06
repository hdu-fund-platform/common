package logs

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	// test default logger with outer functions
	ctx := context.TODO()
	Debug(ctx, "debug: %d", 1)
	Info(ctx, "info: %d", 2)
	Warn(ctx, "warn: %d", 3)
	Error(ctx, "error: %d", 4)
}
