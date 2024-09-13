// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure for table words.
type Words struct {
	Id        uint        `json:"id"        orm:"id"         ` // Word ID
	Content   string      `json:"content"   orm:"content"    ` // Word content
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // Created Time
}
