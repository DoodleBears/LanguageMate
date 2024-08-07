package word

import (
	"context"

	v1 "backend/api/word/v1"
	"backend/internal/dao"
)

func (c *ControllerV1) IndexWord(ctx context.Context, req *v1.IndexWordReq) (res *v1.IndexWordRes, err error) {
	res = &v1.IndexWordRes{}
	err = dao.Words.Ctx(ctx).Scan(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
