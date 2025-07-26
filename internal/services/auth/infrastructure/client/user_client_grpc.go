package authclientgrpc

import (
	"context"

	userpb "github.com/Noname2812/go-ecommerce-backend-api/internal/common/protogen/user"
	grpcserver "github.com/Noname2812/go-ecommerce-backend-api/pkg/grpc"
)

type UserGRPCClient struct {
	client userpb.UserServiceClient
}

// NewUserGRPCClient creates a new gRPC client with modern best practices
func NewUserGRPCClient(manager *grpcserver.GRPCServerManager) *UserGRPCClient {
	userClient, err := manager.AddClient("user-client", "127.0.0.1:4001")
	if err != nil {
		panic("connect user grpc failed !")
	}
	conn := userClient.GetConnection()
	return &UserGRPCClient{
		client: userpb.NewUserServiceClient(conn),
	}
}

func (c *UserGRPCClient) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	return c.client.CreateUser(ctx, req)
}
