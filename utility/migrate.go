package utility

import (
	"errors"
	"io/fs"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func migrateConnect(fsys fs.FS) *migrate.Migrate {
	ctx := gctx.GetInitCtx()
	link, err := g.Cfg().Get(ctx, "database.default.link")
	if err != nil {
		panic(err)
	}

	// GoFrame link format: "mysql:user:pass@tcp(host:port)/db"
	// golang-migrate URL:  "mysql://user:pass@tcp(host:port)/db"
	dsn := strings.Replace(link.String(), "mysql:", "mysql://", 1) + "?multiStatements=true"

	d, err := iofs.New(fsys, "manifest/sql")
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dsn)
	if err != nil {
		panic(err)
	}
	return m
}

func MigrateUp(fsys fs.FS) {
	m := migrateConnect(fsys)
	defer m.Close()
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}

func MigrateDown(fsys fs.FS) {
	m := migrateConnect(fsys)
	defer m.Close()
	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}
