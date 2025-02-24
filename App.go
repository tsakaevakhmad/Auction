package main

import (
	"Auction/controllers"
	"Auction/handlers/category"
	"Auction/services/dbcontext"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var basePath = "/api/v1/"
var Module = fx.Options(
	fx.Provide(
		//Services
		dbcontext.NewPgContext,
		category.NewCreateCategoryHandler,
		category.NewGetCategoryHandler,
		//Controllers
		controllers.NewCategoryControler,
	),
)

var server = fx.Invoke(func(category *controllers.CategoryControler) {
	router := gin.Default()
	router.POST(basePath+"category/create", category.Create)
	router.POST(basePath+"category/get", category.Get)
	router.Run(":8080")
})
