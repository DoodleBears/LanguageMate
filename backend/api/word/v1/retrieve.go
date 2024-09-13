package v1

import (
	v1 "backend/api/sentence/v1"

	"github.com/gogf/gf/v2/frame/g"
)

type Word struct {
	Id      int    `json:"id" des:"Word id" eg:"1"`
	Content string `json:"content" des:"Word content text" eg:"abandon"`
}

// parameters should use uppercase
type RetrieveWordReq struct {
	g.Meta `path:"/:id" tags:"Word" method:"get" sum:"Get a word"`
}
type RetrieveWordRes struct {
	Content   string          `json:"content" des:"Word content text" eg:"abandon"`
	Sentences v1.SentenceList `json:"sentences" des:"Related sentences"`
}
