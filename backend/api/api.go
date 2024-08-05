package api

import "github.com/gogf/gf/v2/frame/g"

type CommonRes struct {
	g.Meta  `mime:"application/json"`
	Code    int         `json:"code" des:"Status Code"`
	Message string      `json:"message" des:"Status Message"`
	Data    interface{} `json:"data"`
}
