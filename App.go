package main

import (
	"Auction/controllers"
	_ "Auction/docs"
	"Auction/handlers/category"
	"Auction/services/dbcontext"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

var BasePath = "/api/v1/"
var module = fx.Options(
	fx.Provide(
		//Services
		dbcontext.NewPgContext,
		category.NewCreateCategoryHandler,
		category.NewGetCategoryHandler,
		category.NewUpdateCategoryHandler,
		category.NewDeleteCategoryHandler,
		//Controllers
		controllers.NewCategoryControler,
	),
)

var server = fx.Invoke(func(category *controllers.CategoryControler) {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST(BasePath+"category/create", category.Create)
	router.POST(BasePath+"category/getall", category.GetAll)
	router.PUT(BasePath+"category/update", category.Update)
	router.DELETE(BasePath+"category/delete/:id", category.Delete)
	router.Run(":8080")
})
