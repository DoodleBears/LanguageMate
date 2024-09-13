package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type User struct {
	g.Meta    `orm:"table:users"`
	Uid       string `json:"uid" eg:"1000000000"`
	AvatarUrl string `json:"avatarUrl" eg:"default"`
	Username  string `json:"username" eg:"chair"`
	Email     string `json:"email" eg:"1@email.com"`
	Nickname  string `json:"nickname" eg:"chair"`
}

type GetCurrentUserReq struct {
	g.Meta `path:"/current" tags:"User" method:"get" sum:"Get Current User"`
}

type GetCurrentUserRes *User
