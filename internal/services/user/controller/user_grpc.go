package usercontroller

import (
	userpb "github.com/Noname2812/go-ecommerce-backend-api/internal/common/protogen/user"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/initialize"
	userservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/service"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/grpcserver"
	"google.golang.org/grpc"
)

// InitUserGrpcServer initializes and returns a GRPCServer for user service
func InitUserGrpcServer(appContainer *initialize.AppContainer, addr string) *grpcserver.GRPCServer {

	// Register gRPC service
	userServiceServer := userservice.NewUserServiceServer()
	registerFunc := func(server *grpc.Server) {
		userpb.RegisterUserServiceServer(server, userServiceServer) // generated proto Register
	}

	// Create GRPC server with the registerFunc
	return grpcserver.NewGRPCServer(addr, registerFunc)
}
