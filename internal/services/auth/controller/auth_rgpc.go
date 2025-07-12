package authcontroller

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func InitAuthGrpcServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	// register grpc service

	gRPCServer := grpc.NewServer()
	log.Println("gRPC server listening on", s.addr)
	return gRPCServer.Serve(lis)
}
