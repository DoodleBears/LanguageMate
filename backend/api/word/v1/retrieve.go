package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// parameters should use uppercase
type RetrieveWordReq struct {
	g.Meta `path:"/:id" tags:"Word" method:"get" sum:"Get a word"`
}
type RetrieveWordRes string
