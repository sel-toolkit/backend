package main

import (
	"backend/env"
	_ "backend/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/v2/os/gctx"

	"backend/internal/cmd"
)

func main() {
	env.InitEnv()

	// Explicitly set the configuration file path to config.yaml
	g.Cfg().SetFileName("manifest/config/config.yaml")

	cmd.Main.Run(gctx.GetInitCtx())
}
