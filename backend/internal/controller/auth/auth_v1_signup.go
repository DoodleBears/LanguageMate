package auth

import (
	"context"

	v1 "backend/api/auth/v1"
	"backend/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error) {
	r := g.RequestFromCtx(ctx)
	user, err := service.User().CreateUser(ctx, req)
	if err != nil {
		r.Response.Status = 500
		return nil, err
	}
	res = &v1.SignupRes{}
	SessionId := service.Session().GenerateSessionId(r)
	r.Session.Set("uid", user.Uid)
	r.Cookie.SetSessionId(SessionId)
	r.Response.Status = 201
	return res, nil
}
