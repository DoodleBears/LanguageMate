package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type Sentence struct {
	WordId  uint   `json:"word_id" v:"required" des:"Related word id" eg:"1"`
	Content string `json:"content" des:"Word content text" eg:"We need to abandon the car."`
}
type SentenceList []*Sentence

// parameters should use uppercase
type IndexSentenceReq struct {
	g.Meta `path:"/" tags:"Sentence" method:"post" sum:"Get sentences"`
}
type IndexSentenceRes struct {
	Words SentenceList `json:"sentences" des:"Sentence list"`
	Page  int          `json:"page" des:"Current page number"`
	Total int          `json:"total" des:"Total page number"`
}
