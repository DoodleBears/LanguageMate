// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package word

import (
	"context"

	"backend/api/word/v1"
)

type IWordV1 interface {
	CreateWord(ctx context.Context, req *v1.CreateWordReq) (res *v1.CreateWordRes, err error)
	RetrieveWord(ctx context.Context, req *v1.RetrieveWordReq) (res *v1.RetrieveWordRes, err error)
}
