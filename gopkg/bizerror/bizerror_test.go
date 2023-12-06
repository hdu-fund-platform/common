package bizerror

import (
	"context"
	"github.com/hdu-fund-platform/common/gopkg/logs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBizError(t *testing.T) {
	ch := make(chan struct{})
	ctx := context.TODO()
	ctx = logs.CtxAddKVs(ctx, "key", "value")
	err := NewBizError(ctx, 0, "success",
		WithAsyncExecutor(func(ctx context.Context, bizError *BizError) {
			ch <- struct{}{}
		}),
		WithChannelRespOption(200, "success"),
		WithDetailOption("detail msg"),
		WithLogLevelOption(InfoLevel),
		WithStackRows(1),
		WithStackOption(true),
		WithSkipDepthOption(0),
	)
	t.Log("wait for async executor running...")
	<-ch
	t.Log("async execute success")
	assert.EqualValues(t, 0, err.GetCode())
}
