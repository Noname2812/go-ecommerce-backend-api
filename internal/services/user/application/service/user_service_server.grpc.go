package userservice

import (
	userpb "github.com/Noname2812/go-ecommerce-backend-api/internal/common/protogen/user"
)

type userServiceServer struct {
	userpb.UnimplementedUserServiceServer
}

func NewUserServiceServer() userpb.UserServiceServer {
	return &userServiceServer{}
}
