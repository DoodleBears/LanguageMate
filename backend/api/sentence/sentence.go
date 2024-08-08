// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sentence

import (
	"context"

	"backend/api/sentence/v1"
)

type ISentenceV1 interface {
	CreateSentence(ctx context.Context, req *v1.CreateSentenceReq) (res *v1.CreateSentenceRes, err error)
	RetrieveSentence(ctx context.Context, req *v1.RetrieveSentenceReq) (res *v1.RetrieveSentenceRes, err error)
}
