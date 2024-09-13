package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta     `path:"/login" tags:"Auth" method:"post" sum:"User Login"`
	Username   string `json:"username" v:"required-without:Email" des:"User Name" eg:"admin"`
	Email      string `json:"email" v:"required-without:Username" des:"User Email" eg:"admin@gmail.com"`
	Password   string `json:"password" v:"required" des:"Password" eg:"adminPassword"`
	RememberMe bool   `json:"remember_me" v:"required" des:"Remember me flag" eg:"true"`
}

type LoginRes struct {
	SessionId string ` json:"session_id" des:"Session ID"`
}
