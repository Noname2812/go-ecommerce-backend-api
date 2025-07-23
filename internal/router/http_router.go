package router

import (
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/initialize"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/middlewares"
	authcontroller "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/controller"
	usercontroller "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/controller"
)

func InitHttpRouter(config *setting.Config, container *initialize.AppContainer) *gin.Engine {
	var r *gin.Engine
	if config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	v := validator.New()
	// middlewares
	r.Use(middlewares.CORS)
	r.Use(middlewares.ValidatorMiddleware(v))
	r.Use(middlewares.RecoveryWithLogger(container.Logger))
	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/checkStatus") // tracking monitor
	}
	{
		usercontroller.InitUserRouter(MainGroup, container)
		authcontroller.InitAuthRouter(MainGroup, container)
	}

	return r
}
