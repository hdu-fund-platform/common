package logs

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func init() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	config.EncoderConfig = encoderConfig
	l, _ := config.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(2))
	Setup(&ZapLog{log: l})
}

type ZapLog struct {
	log *zap.Logger
}

func (z *ZapLog) Debug(ctx context.Context, format string, args ...any) {
	z.wrapCtx(ctx).Debugf(format, args...)
}

func (z *ZapLog) Info(ctx context.Context, format string, args ...any) {
	z.wrapCtx(ctx).Infof(format, args...)
}

func (z *ZapLog) Warn(ctx context.Context, format string, args ...any) {
	z.wrapCtx(ctx).Warnf(format, args...)
}

func (z *ZapLog) Error(ctx context.Context, format string, args ...any) {
	z.wrapCtx(ctx).Errorf(format, args...)
}

func (z *ZapLog) wrapCtx(ctx context.Context) *zap.SugaredLogger {
	kvs := GetAllKVs(ctx)
	fields := make([]zap.Field, 0, (len(kvs)+1)/2)
	for i := 0; i < len(kvs); i += 2 {
		key, ok := kvs[i].(string)
		if !ok {
			continue
		}
		if i+1 >= len(kvs) {
			break
		}
		value := kvs[i+1]
		fields = append(fields, zap.Any(key, value))
	}
	s := z.log.With(fields...).Sugar()
	return s
}
