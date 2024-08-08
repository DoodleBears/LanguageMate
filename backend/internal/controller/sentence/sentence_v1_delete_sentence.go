package sentence

import (
	"context"

	v1 "backend/api/sentence/v1"
	"backend/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) DeleteSentence(ctx context.Context, req *v1.DeleteSentenceReq) (res *v1.DeleteSentenceRes, err error) {
	r := g.RequestFromCtx(ctx)
	uid := r.Session.MustGet("uid").String()
	id := g.RequestFromCtx(ctx).GetRouter("id").Int()
	// User can only delete owned
	dao.Sentences.Ctx(ctx).Where(g.Map{
		"user_uid": uid,
		"id":       id,
	}).Delete()
	return
}
