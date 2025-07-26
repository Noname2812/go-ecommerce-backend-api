package transportationcontroller

import (
	"github.com/Noname2812/go-ecommerce-backend-api/internal/initialize"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/middlewares"
	transportationwire "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/wire"
	"github.com/gin-gonic/gin"
)

func InitTransportationRouter(rg *gin.RouterGroup, container *initialize.AppContainer) {
	transportationQueryHandler := transportationwire.InitTransportationQueryHandler(container.DB, container.RedisClient, container.LocalCache, container.Logger)
	// public router
	transportationRouterPublic := rg.Group("/transportation")
	{
		transportationRouterPublic.GET("/search-trips", transportationQueryHandler.GetListTrips)
	}

	// private router
	transportationRouterPrivate := rg.Group("/transportation")
	transportationRouterPrivate.Use(middlewares.AuthenMiddleware())
	// transportationRouterPrivate.Use(middlewares.NewRateLimiter().UserAndPrivateRateLimiter())
	// transportationRouterPrivate.Use(Authen())
	// transportationRouterPrivate.Use(Permission())
	// {
	// 	// userRouterPrivate.GET("/get_info/:id", userQueryHandler.GetUserProfile)
	// 	// userRouterPrivate.POST("/two-factor/setup", account.TwoFA.SetupTwoFactorAuth)
	// 	// userRouterPrivate.POST("/two-factor/verify", account.TwoFA.VerifyTwoFactorAuth)
	// }
}
