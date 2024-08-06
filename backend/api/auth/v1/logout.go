package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"Auth" method:"post" sum:"User Logout"`
}
type LogoutRes struct {
}
