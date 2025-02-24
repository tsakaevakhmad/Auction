package main

import (
	"Auction/controllers"
	"Auction/services/dbcontext"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		Module,
		fx.Invoke(func(pgdb *dbcontext.PgContext) {
			pgdb.Migrate()
		}),
		fx.Invoke(func(category *controllers.CategoryControler) {
			router := gin.Default()
			router.POST("/api/v1/category/create", category.Create)
			router.Run(":8080")
		}),
	).Run()
}
