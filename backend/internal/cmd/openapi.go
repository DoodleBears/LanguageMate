package cmd

import (
	"backend/api"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Openapi = gcmd.Command{
		Name:  "openapi",
		Usage: "openapi",
		Brief: "set openapi files for development",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// OpenAPI related setting, should be deleted in production
			s := g.Server()
			openapi := s.GetOpenApi()
			openapi.Config.CommonResponse = api.CommonRes{}
			openapi.Config.CommonResponseDataField = `Data`
			openapi.Info = goai.Info{
				Title:       "LanguageMate API Reference",
				Description: "LanguageMate",
			}
			return nil
		},
	}
)
