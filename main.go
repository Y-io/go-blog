package main

import (
	_ "go-blog/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"go-blog/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
