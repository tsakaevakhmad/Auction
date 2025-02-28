package main

import (
	"Auction/controllers"
	_ "Auction/docs"
	"Auction/handlers/category"
	"Auction/services/auth"
	"Auction/services/dbcontext"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	_ "github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	_ "time"
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
		auth.NewPasskeyService,
		//Controllers
		controllers.NewCategoryControler,
	),
)

var server = fx.Invoke(func(category *controllers.CategoryControler, auth *auth.PassKeyService) {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Разрешенные источники
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))
	store := cookie.NewStore([]byte("secret")) // Replace "secret" with your secret key
	router.Use(sessions.Sessions("session_name", store))
	router.POST("/register/begin", auth.BeginRegistration)
	router.POST("/register/finish", auth.FinishRegistration)
	router.POST("/login/begin", auth.BeginLogin)
	router.POST("/login/finish", auth.FinishLogin)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST(BasePath+"category/create", category.Create)
	router.POST(BasePath+"category/getall", category.GetAll)
	router.PUT(BasePath+"category/update", category.Update)
	router.DELETE(BasePath+"category/delete/:id", category.Delete)
	router.Run(":8080")
})
