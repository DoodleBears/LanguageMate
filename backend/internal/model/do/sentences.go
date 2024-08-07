// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Sentences is the golang structure of table sentences for DAO operations like Where/Data.
type Sentences struct {
	g.Meta    `orm:"table:sentences, do:true"`
	UserUid   interface{} // User Unique ID
	WordId    interface{} // Word ID
	Content   interface{} // Sentence content
	CreatedAt *gtime.Time // Created Time
	DeletedAt *gtime.Time // Soft deleted Time
}
