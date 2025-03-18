package main

import (
	"Auction/controllers"
	_ "Auction/docs"
	"Auction/domain/configurations"
	"Auction/handlers/category"
	"Auction/services/Configuration"
	"Auction/services/auth"
	"Auction/services/dbcontext"
	"fmt"
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
		auth.NewJWTServices,
		//Controllers
		controllers.NewCategoryControler,
	),
)

var server = fx.Invoke(func(category *controllers.CategoryControler, auth *auth.PassKeyService, jwtServices *auth.JWTServices) {
	var cfg *configurations.MainConfig
	Configuration.ReadFile(&cfg)
	router := gin.Default()
	additionalRouterFunctions(router, cfg)
	authorizedAdmin := router.Group(BasePath)
	router.POST("/register/begin/:username", auth.BeginRegistration)
	router.POST("/register/finish/:username", auth.FinishRegistration)
	router.POST("/login/begin/:username", auth.BeginLogin)
	router.POST("/login/finish/:username", auth.FinishLogin)

	authorizedAdmin.Use(jwtServices.AuthMiddleware("admin"))
	{
		authorizedAdmin.POST("category/create", category.Create)
		authorizedAdmin.PUT("category/update", category.Update)
		authorizedAdmin.DELETE("category/delete/:id", category.Delete)
		router.POST(BasePath+"category/getall", category.GetAll)
	}
	router.Run(fmt.Sprint(":", cfg.Server.Port))
})

func additionalRouterFunctions(router *gin.Engine, cfg *configurations.MainConfig) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Server.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Domain:   cfg.Server.Domain,
	})
	router.Use(sessions.Sessions("passKey", store))
	router.Static("/static", "./static")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
