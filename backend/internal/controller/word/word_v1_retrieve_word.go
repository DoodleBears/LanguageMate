package word

import (
	"context"

	sentencev1 "backend/api/sentence/v1"
	v1 "backend/api/word/v1"
	"backend/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) RetrieveWord(ctx context.Context, req *v1.RetrieveWordReq) (res *v1.RetrieveWordRes, err error) {
	res = &v1.RetrieveWordRes{}
	var list sentencev1.SentenceList
	uid := g.RequestFromCtx(ctx).Session.MustGet("uid").String()
	result := &v1.Word{}
	id := g.RequestFromCtx(ctx).GetRouter("id").Int()
	err = dao.Words.Ctx(ctx).Where("id", id).Scan(&result)
	if err != nil {
		return nil, err
	}
	res.Content = result.Content
	err = dao.Sentences.Ctx(ctx).Where(g.Map{
		"user_uid": uid,
		"word_id":  id,
	}).Scan(&list)

	if err != nil {
		return nil, err
	}
	res.Sentences = list
	return
}
