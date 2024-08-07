package word

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/word/v1"
)

func (c *ControllerV1) RetrieveWord(ctx context.Context, req *v1.RetrieveWordReq) (res *v1.RetrieveWordRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
