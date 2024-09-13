// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta        `orm:"table:users, do:true"`
	Uid           interface{} // User Unique ID
	AvatarUrl     interface{} // User Avatar
	Username      interface{} // User Name
	Email         interface{} // User Email
	EmailVerified interface{} // User Email Verification Flag
	Password      interface{} // User Password
	Nickname      interface{} // User Nickname
	CreatedAt     *gtime.Time // Created Time
	UpdatedAt     *gtime.Time // Updated Time
	DeletedAt     *gtime.Time // Soft deleted Time
	LastLogin     *gtime.Time // Last Login Time
	LoginAttempts interface{} // User Login Attempts
	Lock          interface{} // User Login Lock
	Banned        interface{} // User Banned control
}
