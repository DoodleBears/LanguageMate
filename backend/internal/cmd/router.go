package cmd

import (
	"backend/internal/controller/auth"
	"backend/internal/controller/user"
	"backend/internal/controller/word"
	"backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Router = gcmd.Command{
		Name:  "router",
		Usage: "router",
		Brief: "register route",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Group("/auth", func(group *ghttp.RouterGroup) {
					// Add ctx never done since some api are using goroutine to send email
					// Could be pref in the future
					group.Middleware(service.Middleware().NeverDoneCtx)
					group.Bind(
						auth.NewV1(),
					)
				})
				group.Group("/users", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						user.NewV1(),
					)
				})
				group.Group("/words", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						word.NewV1(),
					)
				})
			})
			return nil
		},
	}
)
