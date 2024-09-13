package sentence

import (
	"context"

	v1 "backend/api/sentence/v1"
	"backend/internal/dao"
	"backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) CreateSentence(ctx context.Context, req *v1.CreateSentenceReq) (res *v1.CreateSentenceRes, err error) {
	r := g.RequestFromCtx(ctx)
	sentence := &entity.Sentences{}
	sentence.Content = req.Content
	sentence.UserUid = r.Session.MustGet("uid").String()
	sentence.WordId = req.WordId
	_, err = dao.Sentences.Ctx(ctx).Data(sentence).Save()
	if err != nil {
		return nil, err
	}
	return
}
