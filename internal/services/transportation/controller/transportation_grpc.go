package transportationcontroller

import (
	transportationpb "github.com/Noname2812/go-ecommerce-backend-api/internal/common/protogen/transportation"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/initialize"
	transportationwire "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/wire"
	grpcserver "github.com/Noname2812/go-ecommerce-backend-api/pkg/grpc"

	"google.golang.org/grpc"
)

// InitTransportationGrpcServer initializes and returns a GRPCServer for transportation service
func InitTransportationGrpcServer(appContainer *initialize.AppContainer, addr string) *grpcserver.GRPCServer {

	// Register gRPC service
	transportationServiceServer := transportationwire.InitTransportationServiceServer(appContainer.DB, appContainer.Logger, appContainer.RedisClient)
	registerFunc := func(server *grpc.Server) {
		transportationpb.RegisterTransportationServiceServer(server, transportationServiceServer) // generated proto Register
	}

	// Create GRPC server with the registerFunc
	return grpcserver.NewGRPCServer(addr, registerFunc)
}
