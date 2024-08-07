package word

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/word/v1"
)

func (c *ControllerV1) CreateWord(ctx context.Context, req *v1.CreateWordReq) (res *v1.CreateWordRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
