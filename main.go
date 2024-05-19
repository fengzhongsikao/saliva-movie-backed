package main

import (
	_ "saliva-movie/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"saliva-movie/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
