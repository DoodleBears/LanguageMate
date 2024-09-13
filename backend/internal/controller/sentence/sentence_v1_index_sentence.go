package sentence

import (
	"context"

	v1 "backend/api/sentence/v1"
	"backend/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) IndexSentence(ctx context.Context, req *v1.IndexSentenceReq) (res *v1.IndexSentenceRes, err error) {
	r := g.RequestFromCtx(ctx)
	res = &v1.IndexSentenceRes{}
	var list v1.SentenceList
	total := 0
	if req.Page <= 0 || req.PerPage <= 0 {
		err = dao.Sentences.Ctx(ctx).Where("user_uid", r.Session.MustGet("uid").String()).Scan(&list)
	} else {
		err = dao.Sentences.Ctx(ctx).Where("user_uid", r.Session.MustGet("uid")).OmitEmpty().
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
	res.Sentences = list
	return res, nil
}
