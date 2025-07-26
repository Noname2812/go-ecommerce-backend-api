package router

import (
	"fmt"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/initialize"
	usercontroller "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/controller"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
)

func InitGrpcServer(config *setting.Config, appContainer *initialize.AppContainer) {
	// ================== Server =================== //
	// User server
	userGrpcAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.GRPC.UserServicePort)
	userServer := usercontroller.InitUserGrpcServer(appContainer, userGrpcAddr)
	appContainer.GRPCServerManager.AddServer("user-grpc-server", userServer)

	go appContainer.GRPCServerManager.StartAllServers()

}
