package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// parameters should use uppercase
type DeleteSentenceReq struct {
	g.Meta `path:"/:id" tags:"Sentence" method:"delete" sum:"Delete a sentence"`
}
type DeleteSentenceRes struct{}
