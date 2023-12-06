package bizerror

import (
	"bytes"
	"context"
	"fmt"
	"github.com/hdu-fund-platform/common/gopkg"
	"github.com/hdu-fund-platform/common/gopkg/logs"
	"runtime"
)

// BizError 自定义Error类型(实现了go内嵌error接口)
// 特性:
//  1. 包含规范的服务返回三元组(Code + Msg), 便于封装response
//  2. 符合规范的自动日志打印(NewBizError时打印)
//  3. 堆栈信息
//  4. 其他option拓展见使用说明
type BizError struct {
	code        int64
	msg         string
	level       BizErrLevel
	detail      string
	storeStack  bool
	stack       []byte
	stackRows   int
	depth       int
	channelCode int64                            // 下游错误码
	channelMsg  string                           // 下游错误信息
	asyncFn     func(context.Context, *BizError) // 异步执行函数
}

// BizErrLevel 错误等级, 会影响日志打印时的level
type BizErrLevel int8

// BizErrOption BizError属性设置函数
type BizErrOption func(*BizError)

const (
	// InfoLevel Info级别, 使用logs.CtxInfo打印日志
	InfoLevel BizErrLevel = iota
	// WarnLevel Warn级别, 使用logs.CtxWarn打印日志
	WarnLevel
	// ErrorLevel Warn级别, 使用logs.CtxError打印日志
	ErrorLevel
)

func (e BizError) Error() string {
	errInfo := fmt.Sprintf("code=%d, msg=%s, channelCode=%d, channelMsg=%s, detail=%s",
		e.code, e.msg, e.channelCode, e.channelMsg, e.detail)
	if e.storeStack {
		errInfo = errInfo + "\n" + string(e.stack)
	}
	return errInfo
}

func (e BizError) GetCode() int64 {
	return e.code
}

func (e BizError) GetMsg() string {
	return e.msg
}

func (e BizError) GetDetail() string {
	return e.detail
}

func (e BizError) GetChannelCode() int64 {
	return e.channelCode
}

func (e BizError) GetChannelMsg() string {
	return e.channelMsg
}

func NewBizError(ctx context.Context, code int64, msg string, opts ...BizErrOption) *BizError {
	bizErr := &BizError{
		code:       code,
		msg:        msg,
		level:      ErrorLevel,
		storeStack: true,
		depth:      2,
		stackRows:  10,
	}

	for _, opt := range opts {
		opt(bizErr)
	}

	if bizErr.storeStack {
		bizErr.stack = getStack(bizErr.depth, bizErr.stackRows)
	}

	bizErr.ctxLog(ctx)

	if bizErr.asyncFn != nil {
		gopkg.SafeGo(ctx, func() {
			bizErr.asyncFn(ctx, bizErr)
		})
	}
	return bizErr
}

// WithLogLevelOption 设置日志打印等级, 不设置时默认为ErrorLevel
func WithLogLevelOption(level BizErrLevel) BizErrOption {
	return func(e *BizError) {
		e.level = level
	}
}

// WithDetailOption 设置报错详细信息, 如单号/Uid等参数
func WithDetailOption(format string, v ...interface{}) BizErrOption {
	return func(e *BizError) {
		e.detail = fmt.Sprintf(format, v...)
	}
}

// WithStackOption 设置是否保存函数栈信息, 不设置时默认保存
func WithStackOption(storeStack bool) BizErrOption {
	return func(e *BizError) {
		e.storeStack = storeStack
	}
}

// WithSkipDepthOption 设置跳过的函数栈深度, 当你封装NewBizError时应该设置
func WithSkipDepthOption(skipDepth int) BizErrOption {
	return func(e *BizError) {
		e.depth += skipDepth
	}
}

// WithChannelRespOption 设置下游返回的错误码/消息, 当异常是下游导致的可以设置
func WithChannelRespOption(channelCode int64, channelMsg string) BizErrOption {
	return func(e *BizError) {
		e.channelCode = channelCode
		e.channelMsg = channelMsg
	}
}

// WithAsyncExecutor 产生错误后异步执行器, 如进行上报metrics打点
func WithAsyncExecutor(fn func(context.Context, *BizError)) BizErrOption {
	return func(e *BizError) {
		e.asyncFn = fn
	}
}

// WithStackRows 函数堆栈保存的行数, 默认保存10行
func WithStackRows(stackRows int) BizErrOption {
	return func(e *BizError) {
		if stackRows > 0 {
			e.stackRows = stackRows
		}
	}
}

func (e BizError) ctxLog(ctx context.Context) {
	switch e.level {
	case InfoLevel:
		logs.Info(ctx, "%s", e.Error())
	case WarnLevel:
		logs.Warn(ctx, "%s", e.Error())
	case ErrorLevel:
		logs.Error(ctx, "%s", e.Error())
	default:
		logs.Error(ctx, "%s", e.Error())
	}
}

// getStack returns a nicely formated stack frame, skipping skip frames
func getStack(skip, rows int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	for i := skip; i-skip < rows; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		_, _ = fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.Bytes()
}
