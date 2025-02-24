package main

import (
	"Auction/services/dbcontext"
	"fmt"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	fx.New(
		Module,
		fx.Invoke(func(pgdb *dbcontext.PgContext) {
			pgdb.Migrate()
		}),
		fx.Invoke(func(mux *http.ServeMux) {
			fmt.Println("Server started on :8080")
			http.ListenAndServe(":8080", mux)
		}),
	).Run()
}
