package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// parameters should use uppercase
type CreateWordReq struct {
	g.Meta  `path:"/" tags:"Word" method:"post" sum:"Create a word"`
	Content string `json:"content" v:"required|length:1,45" des:"Word content text" eg:"abandon"`
}
type CreateWordRes struct {
}
