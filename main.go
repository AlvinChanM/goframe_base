package main

import (
	_ "github.com/AlvinChanM/goframe_base/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/AlvinChanM/goframe_base/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
