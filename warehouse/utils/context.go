package warehouseutils

import (
	"context"
)

type uploadIDContextKey struct{}

func CtxWithUploadID(ctx context.Context, uid int64) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func UploadIDFromCtx(ctx context.Context) (int64, bool) { _ = "STUB: not implemented"; return 0, false }
