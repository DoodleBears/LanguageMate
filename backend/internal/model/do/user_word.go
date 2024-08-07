// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserWord is the golang structure of table user_word for DAO operations like Where/Data.
type UserWord struct {
	g.Meta    `orm:"table:user_word, do:true"`
	UserUid   interface{} // User Unique ID
	WordId    interface{} // Word ID
	CreatedAt *gtime.Time // Created Time
	DeletedAt *gtime.Time // Soft deleted Time
}
