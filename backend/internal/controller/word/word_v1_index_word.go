package word

import (
	"context"

	v1 "backend/api/word/v1"
	"backend/internal/dao"
)

func (c *ControllerV1) IndexWord(ctx context.Context, req *v1.IndexWordReq) (res *v1.IndexWordRes, err error) {
	res = &v1.IndexWordRes{}
	var list v1.WordsList
	total := 0
	if req.Page <= 0 || req.PerPage <= 0 {
		err = dao.Words.Ctx(ctx).Scan(&list)
	} else {
		err = dao.Words.Ctx(ctx).OmitEmpty().
			Order("id", "asc").
			Limit((req.Page-1)*req.PerPage, req.PerPage).
			ScanAndCount(&list, &total, true)
	}
	if err != nil {
		return nil, err
	}
	if req.Page > 0 && req.PerPage > 0 {
		res.Page = req.Page
		remainder := total % req.PerPage
		brefPage := total / req.PerPage
		if remainder == 0 {
			res.Total = brefPage
		} else {
			res.Total = brefPage + 1
		}
	}
	res.Words = list
	return res, nil
}
