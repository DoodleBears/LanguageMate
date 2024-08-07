package word

import (
	"context"

	v1 "backend/api/word/v1"
	"backend/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) RetrieveWord(ctx context.Context, req *v1.RetrieveWordReq) (res *v1.RetrieveWordRes, err error) {
	res = &v1.RetrieveWordRes{}
	id := g.RequestFromCtx(ctx).GetRouter("id").Int()
	err = dao.Words.Ctx(ctx).Where("id", id).Scan(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
