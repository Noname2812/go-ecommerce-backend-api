package grpcserver

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	server   *grpc.Server
	addr     string
	register func(*grpc.Server)
}

func NewGRPCServer(addr string, register func(*grpc.Server)) *GRPCServer {
	return &GRPCServer{
		server:   grpc.NewServer(),
		addr:     addr,
		register: register,
	}
}

func (s *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	// Register gRPC handlers
	if s.register != nil {
		s.register(s.server)
	}

	log.Println("gRPC server listening on", s.addr)
	return s.server.Serve(lis)
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
	log.Println("gRPC server stopped gracefully")
}

func (s *GRPCServer) Addr() string {
	return s.addr
}
