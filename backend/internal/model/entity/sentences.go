// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Sentences is the golang structure for table sentences.
type Sentences struct {
	UserUid   string      `json:"userUid"   orm:"user_uid"   ` // User Unique ID
	WordId    uint        `json:"wordId"    orm:"word_id"    ` // Word ID
	Content   string      `json:"content"   orm:"content"    ` // Sentence content
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // Created Time
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // Soft deleted Time
}
