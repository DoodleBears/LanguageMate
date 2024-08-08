package sentence

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/sentence/v1"
)

func (c *ControllerV1) RetrieveSentence(ctx context.Context, req *v1.RetrieveSentenceReq) (res *v1.RetrieveSentenceRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
