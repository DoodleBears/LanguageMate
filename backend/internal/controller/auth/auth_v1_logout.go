package auth

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "backend/api/auth/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	r := g.RequestFromCtx(ctx)
	r.Session.RemoveAll()
	return
}
