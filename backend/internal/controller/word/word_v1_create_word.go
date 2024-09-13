package word

import (
	"context"

	v1 "backend/api/word/v1"
	"backend/internal/dao"
)

func (c *ControllerV1) CreateWord(ctx context.Context, req *v1.CreateWordReq) (res *v1.CreateWordRes, err error) {
	result, err := dao.Words.Ctx(ctx).Data(req).Save()
	if err != nil {
		return nil, err
	}
	res = &v1.CreateWordRes{}
	res.Id, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return res, nil
}
