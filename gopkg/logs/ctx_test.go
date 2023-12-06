package logs

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CtxAddKVs(t *testing.T) {
	t.Run("Test adding kvs not even and ignore this operate", func(t *testing.T) {
		ctx := context.TODO()
		ctx = CtxAddKVs(ctx, "key")
		kvs := GetAllKVs(ctx)
		assert.Len(t, kvs, 0)
	})

	t.Run("Test adding multi even kvs and success", func(t *testing.T) {
		ctx := context.TODO()
		ctx = CtxAddKVs(ctx, "k1", "v1")
		ctx = CtxAddKVs(ctx, "k2", "v2")
		kvs := GetAllKVs(ctx)
		want := []interface{}{"k1", "v1", "k2", "v2"}
		assert.Equal(t, want, kvs)
	})
}

func TestCtxLog(t *testing.T) {
	ctx := context.TODO()
	ctx = CtxAddKVs(ctx, "k", "v")
	Debug(ctx, "test")
}
