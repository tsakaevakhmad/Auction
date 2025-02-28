package main

import (
	"Auction/services/dbcontext"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Invoke(func(pgdb *dbcontext.PgContext) {
			pgdb.Migrate()
		}),
		module,
		server,
	).Run()
}
