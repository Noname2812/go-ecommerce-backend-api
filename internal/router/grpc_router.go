package router

import (
	"fmt"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/initialize"
	usercontroller "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/controller"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/grpcserver"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
)

func InitGrpcServer(config *setting.Config, appContainer *initialize.AppContainer) {
	servers := []*grpcserver.GRPCServer{}
	// user grpc
	userGrpcAddr := fmt.Sprintf("localhost:%d", config.GRPC.UserServicePort)
	userServer := usercontroller.InitUserGrpcServer(appContainer, userGrpcAddr)
	servers = append(servers, userServer)

	// Start all gRPC servers concurrently
	for _, srv := range servers {
		go func(s *grpcserver.GRPCServer) {
			appContainer.Logger.Info("Starting gRPC server",
				zap.String("address", s.Addr()))

			if err := s.Start(); err != nil {
				appContainer.Logger.Fatal("Failed to start gRPC server",
					zap.String("address", s.Addr()),
					zap.Error(err))
			}
		}(srv)
	}
}
