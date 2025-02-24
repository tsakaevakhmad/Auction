package main

import (
	"Auction/controllers"
	"Auction/handlers/category"
	"Auction/services/dbcontext"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		//Services
		dbcontext.NewPgContext,
		category.NewCreateCategoryHandler,

		//Controllers
		controllers.NewCategoryControler,
	),
)
