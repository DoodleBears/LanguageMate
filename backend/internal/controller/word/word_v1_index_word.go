package word

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/word/v1"
)

func (c *ControllerV1) IndexWord(ctx context.Context, req *v1.IndexWordReq) (res *v1.IndexWordRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
