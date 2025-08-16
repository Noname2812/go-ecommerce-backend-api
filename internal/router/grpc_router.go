package router

import (
	"fmt"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/initialize"
	transportationcontroller "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/controller"
	usercontroller "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/controller"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
)

const (
	USER_SERVICE_NAME           = "user-grpc-server"
	TRANSPORTATION_SERVICE_NAME = "transportation-grpc-server"
)

func InitGrpcServer(config *setting.Config, appContainer *initialize.AppContainer) {
	// ================== Server =================== //
	// User server
	userGrpcAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.GRPC.UserServicePort)
	userServer := usercontroller.InitUserGrpcServer(appContainer, userGrpcAddr)
	appContainer.GRPCServerManager.AddServer(USER_SERVICE_NAME, userServer)

	// Transportation server
	transportationGrpcAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.GRPC.TransportationServicePort)
	transportationServer := transportationcontroller.InitTransportationGrpcServer(appContainer, transportationGrpcAddr)
	appContainer.GRPCServerManager.AddServer(TRANSPORTATION_SERVICE_NAME, transportationServer)

	go appContainer.GRPCServerManager.StartAllServers()

}
