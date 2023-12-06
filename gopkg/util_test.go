package gopkg

import (
	"context"
	"testing"
	"time"
)

func Test_SafeGo(t *testing.T) {
	SafeGo(context.TODO(), func() {
		panic("hello")
	})

	time.Sleep(time.Millisecond * 10)
}

func Test_SafeGoWithCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	SafeGoWithCtx(ctx, func(ctx context.Context) {
		select {
		case <-ctx.Done():
			panic("hello")
		}
	})

	cancel()
	time.Sleep(time.Millisecond * 10)
}
