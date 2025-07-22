package grpcserver

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type GRPCClient struct {
	conn *grpc.ClientConn
	addr string
}

// ClientConfig contains configuration for gRPC client
type ClientConfig struct {
	Addr              string
	MaxRetryAttempts  int
	ConnectionTimeout time.Duration
	KeepAlive         time.Duration
	KeepAliveTimeout  time.Duration
}

// DefaultClientConfig returns default configuration for gRPC client
func DefaultClientConfig(addr string) *ClientConfig {
	return &ClientConfig{
		Addr:              addr,
		MaxRetryAttempts:  3,
		ConnectionTimeout: 10 * time.Second,
		KeepAlive:         2 * time.Minute,
		KeepAliveTimeout:  5 * time.Second,
	}
}

// NewGRPCClient creates a new gRPC client with default configuration
func NewGRPCClient(addr string) (*GRPCClient, error) {
	config := DefaultClientConfig(addr)
	return NewGRPCClientWithConfig(config)
}

// NewGRPCClientWithConfig creates a new gRPC client with custom configuration
func NewGRPCClientWithConfig(config *ClientConfig) (*GRPCClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.ConnectionTimeout)
	defer cancel()

	// Connection options
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), // Block until connection is established
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                config.KeepAlive,
			Timeout:             config.KeepAliveTimeout,
			PermitWithoutStream: true,
		}),
	}

	// Establish connection
	conn, err := grpc.DialContext(ctx, config.Addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server at %s: %w", config.Addr, err)
	}

	log.Printf("Successfully connected to gRPC server at %s", config.Addr)

	return &GRPCClient{
		conn: conn,
		addr: config.Addr,
	}, nil
}

// GetConnection returns the underlying gRPC connection
func (c *GRPCClient) GetConnection() *grpc.ClientConn {
	return c.conn
}

// Close closes the gRPC connection
func (c *GRPCClient) Close() error {
	if c.conn != nil {
		log.Printf("Closing gRPC connection to %s", c.addr)
		return c.conn.Close()
	}
	return nil
}

// IsConnected checks if the connection is still active
func (c *GRPCClient) IsConnected() bool {
	if c.conn == nil {
		return false
	}

	state := c.conn.GetState()
	return state == connectivity.Connecting || state == connectivity.Ready
}

// Reconnect attempts to reconnect to the gRPC server
func (c *GRPCClient) Reconnect() error {
	if c.conn != nil {
		c.conn.Close()
	}

	config := DefaultClientConfig(c.addr)
	newClient, err := NewGRPCClientWithConfig(config)
	if err != nil {
		return err
	}

	c.conn = newClient.conn
	return nil
}

// Addr returns the server address
func (c *GRPCClient) Addr() string {
	return c.addr
}
