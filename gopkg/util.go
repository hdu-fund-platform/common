package gopkg

import (
	"context"
	"github.com/hdu-fund-platform/common/gopkg/logs"
	"runtime"
)

const burSize int = 3000

// SafeGo 异步执行并在发生panic后recover&打印堆栈
func SafeGo(ctx context.Context, f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, burSize)
				buf = buf[:runtime.Stack(buf, false)]
				logs.Error(ctx, "SafeGo has panic:%s", string(buf))
			}
		}()
		f()
	}()
}

// SafeGoWithCtx 需要异步执行的函数入参中带有ctx就用这个
func SafeGoWithCtx(ctx context.Context, f func(ctx context.Context)) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, burSize)
				buf = buf[:runtime.Stack(buf, false)]
				logs.Error(ctx, "SafeGo has panic:%s", string(buf))
			}
		}()
		f(ctx)
	}()
}
