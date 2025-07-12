package usercontroller

import (
	"github.com/Noname2812/go-ecommerce-backend-api/internal/initialize"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/middlewares"
	userwire "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/wire"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(rg *gin.RouterGroup, container *initialize.AppContainer) {

	userQueryHandler := userwire.InitUserQueryHandler(container.DB, container.Logger)
	// userCommandHandler := userwire.InitUserCommandHandler(container)

	// public router
	userRouterPublic := rg.Group("/user")
	{
		// userRouterPublic.POST("/register", userCommandHandler.Register)
		userRouterPublic.GET("/:id", userQueryHandler.GetUserDetails)
		// userRouterPublic.POST("/verify_account", account.Login.VerifyOTP)
		// userRouterPublic.POST("/update_pass_register", account.Login.UpdatePasswordRegister)
		// userRouterPublic.POST("/login", account.Login.Login) // login -> YES -> No
	}

	// private router
	userRouterPrivate := rg.Group("/user")
	userRouterPrivate.Use(middlewares.AuthenMiddleware())
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		// userRouterPrivate.GET("/get_info/:id", userQueryHandler.GetUserProfile)
		// userRouterPrivate.POST("/two-factor/setup", account.TwoFA.SetupTwoFactorAuth)
		// userRouterPrivate.POST("/two-factor/verify", account.TwoFA.VerifyTwoFactorAuth)
	}
}
