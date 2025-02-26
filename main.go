package main

import (
	"Auction/services/dbcontext"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		module,
		server,
		fx.Invoke(func(pgdb *dbcontext.PgContext) {
			pgdb.Migrate()
		}),
	).Run()
}
