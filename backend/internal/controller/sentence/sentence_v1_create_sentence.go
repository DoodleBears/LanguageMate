package sentence

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "backend/api/sentence/v1"
)

func (c *ControllerV1) CreateSentence(ctx context.Context, req *v1.CreateSentenceReq) (res *v1.CreateSentenceRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
