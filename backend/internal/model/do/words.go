// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure of table words for DAO operations like Where/Data.
type Words struct {
	g.Meta    `orm:"table:words, do:true"`
	Id        interface{} // Word ID
	Content   interface{} // Word content
	CreatedAt *gtime.Time // Created Time
}
