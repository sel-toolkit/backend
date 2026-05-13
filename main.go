package main

import (
	"backend/env"
	_ "backend/internal/packed"
	"backend/utility"
	"flag"
	"fmt"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/v2/os/gctx"

	"backend/internal/cmd"
)

func main() {
	migrateCmd := flag.String("migrate", "", "database migration command: up / down")
	flag.Parse()

	env.InitEnv()
	g.Cfg().SetFileName("manifest/config/config.yaml")

	switch *migrateCmd {
	case "up":
		utility.MigrateUp(migrationsFS)
		fmt.Println("Migration up completed.")
		return
	case "down":
		utility.MigrateDown(migrationsFS)
		fmt.Println("Migration down completed.")
		return
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
