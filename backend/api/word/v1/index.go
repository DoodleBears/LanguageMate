package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type WordsList []*Word
type IndexWordReq struct {
	g.Meta  `path:"/" tags:"Word" method:"get" sum:"Get words"`
	Page    int `json:"page" des:"Page" in:"query"`
	PerPage int `json:"per_page" des:"Page limit" in:"query"`
}
type IndexWordRes struct {
	Words WordsList `json:"words" des:"Word list"`
	Page  int       `json:"page" des:"Current page number"`
	Total int       `json:"total" des:"Total page number"`
}
