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
	"net/http"
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
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	store := cookie.NewStore([]byte("secret")) // Replace "secret" with your secret key
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600,                 // Время жизни cookie
		HttpOnly: true,                 // Запрещает доступ к cookie через JavaScript
		Secure:   true,                 // Требует HTTPS
		SameSite: http.SameSiteLaxMode, // Разрешает отправку cookie между разными доменами
		Domain:   "localhost",
	})
	router.Use(sessions.Sessions("passKey", store))
	router.Static("/static", "./static")
	router.POST("/register/begin/:username", auth.BeginRegistration)
	router.POST("/register/finish/:username", auth.FinishRegistration)
	router.POST("/login/begin/:username", auth.BeginLogin)
	router.POST("/login/finish/:username", auth.FinishLogin)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST(BasePath+"category/create", category.Create)
	router.POST(BasePath+"category/getall", category.GetAll)
	router.PUT(BasePath+"category/update", category.Update)
	router.DELETE(BasePath+"category/delete/:id", category.Delete)
	router.Run(":8080")
})
