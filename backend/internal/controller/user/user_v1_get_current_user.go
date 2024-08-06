package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "backend/api/user/v1"
	"backend/internal/service"
)

func (c *ControllerV1) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserReq) (res *v1.GetCurrentUserRes, err error) {
	r := g.RequestFromCtx(ctx)
	uid, err := r.Session.Get("uid")
	if err != nil {
		return nil, err
	}
	var user *v1.User
	result, err := service.User().GetUserFromUid(ctx, uid.String())
	gconv.Struct(result, &user)
	res = (*v1.GetCurrentUserRes)(&user)
	return
}
