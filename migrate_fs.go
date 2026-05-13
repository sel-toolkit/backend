package main

import "embed"

//go:embed manifest/sql
var migrationsFS embed.FS
