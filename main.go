package main

import (
	"Auction/services/dbcontext"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		Module,
		fx.Invoke(func(pgdb *dbcontext.PgContext) {
			pgdb.Migrate()
		}),
		server,
	).Run()
}
