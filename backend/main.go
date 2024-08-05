package main

import (
	_ "backend/internal/packed"

	_ "backend/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"backend/internal/cmd"
)

func main() {
	ctx := gctx.GetInitCtx()
	cmd.Router.Run(ctx)
	cmd.Openapi.Run(ctx)
	cmd.Main.Run(ctx)
}
