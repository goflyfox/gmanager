package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "gmanager/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gmanager/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
