package authcontroller

import (
	"github.com/Noname2812/go-ecommerce-backend-api/internal/initialize"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/middlewares"
	authwire "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/wire"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(rg *gin.RouterGroup, container *initialize.AppContainer) {

	authCommandHandler := authwire.InitAuthHttpCommandHandler(container.DB, container.RedisClient, container.Logger, container.KafkaManager)
	// public router
	authRouterPublic := rg.Group("/auth")
	{
		authRouterPublic.POST("/register", authCommandHandler.Register)
		// userRouterPublic.POST("/verify_account", account.Login.VerifyOTP)
		// userRouterPublic.POST("/update_pass_register", account.Login.UpdatePasswordRegister)
		// userRouterPublic.POST("/login", account.Login.Login) // login -> YES -> No
	}

	// private router
	authRouterPrivate := rg.Group("/auth")
	authRouterPrivate.Use(middlewares.AuthenMiddleware())
	// authRouterPrivate.Use(middlewares.NewRateLimiter().UserAndPrivateRateLimiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		// userRouterPrivate.GET("/get_info/:id", userQueryHandler.GetUserProfile)
		// userRouterPrivate.POST("/two-factor/setup", account.TwoFA.SetupTwoFactorAuth)
		// userRouterPrivate.POST("/two-factor/verify", account.TwoFA.VerifyTwoFactorAuth)
	}
}
