// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Uid           string      `json:"uid"           orm:"uid"            ` // User Unique ID
	AvatarUrl     string      `json:"avatarUrl"     orm:"avatar_url"     ` // User Avatar
	Username      string      `json:"username"      orm:"username"       ` // User Name
	Email         string      `json:"email"         orm:"email"          ` // User Email
	EmailVerified bool        `json:"emailVerified" orm:"email_verified" ` // User Email Verification Flag
	Password      string      `json:"password"      orm:"password"       ` // User Password
	Nickname      string      `json:"nickname"      orm:"nickname"       ` // User Nickname
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     ` // Created Time
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     ` // Updated Time
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     ` // Soft deleted Time
	LastLogin     *gtime.Time `json:"lastLogin"     orm:"last_login"     ` // Last Login Time
	LoginAttempts int         `json:"loginAttempts" orm:"login_attempts" ` // User Login Attempts
	Lock          bool        `json:"lock"          orm:"lock"           ` // User Login Lock
	Banned        bool        `json:"banned"        orm:"banned"         ` // User Banned control
}
