package usercontroller

import (
	userpb "github.com/Noname2812/go-ecommerce-backend-api/internal/common/protogen/user"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/initialize"
	userwire "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/wire"
	grpcserver "github.com/Noname2812/go-ecommerce-backend-api/pkg/grpc"

	"google.golang.org/grpc"
)

// InitUserGrpcServer initializes and returns a GRPCServer for user service
func InitUserGrpcServer(appContainer *initialize.AppContainer, addr string) *grpcserver.GRPCServer {

	// Register gRPC service
	userServiceServer := userwire.InitUserServiceServer(appContainer.DB, appContainer.Logger)
	registerFunc := func(server *grpc.Server) {
		userpb.RegisterUserServiceServer(server, userServiceServer) // generated proto Register
	}

	// Create GRPC server with the registerFunc
	return grpcserver.NewGRPCServer(addr, registerFunc)
}
