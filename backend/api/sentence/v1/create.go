package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// parameters should use uppercase
type CreateSentenceReq struct {
	g.Meta  `path:"/" tags:"Sentence" method:"post" sum:"Create a sentence"`
	WordId  uint   `json:"word_id" v:"required" des:"Related word id" eg:"1"`
	Content string `json:"content" v:"required|length:1,200" des:"Sentence content text" eg:"This is a sentence."`
}
type CreateSentenceRes struct{}
