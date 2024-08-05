package cmd

import (
	"backend/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "do basic settings and start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(
				service.Middleware().CORS,
				service.Middleware().ResponseHandler,
			)
			s.EnableAdmin("/system")
			s.BindMiddleware("/system/*", service.Middleware().Auth)
			s.BindMiddleware("/*", service.Middleware().Language)
			s.Run()
			return nil
		},
	}
)
