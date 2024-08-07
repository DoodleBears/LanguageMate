package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IndexWordReq struct {
	g.Meta `path:"/:id" tags:"Word" method:"get" sum:"Get a word"`
}
type IndexWordRes []*Word
