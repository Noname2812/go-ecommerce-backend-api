package grpcserver

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// GRPCServerManager manages both gRPC servers and clients
type GRPCServerManager struct {
	servers map[string]*GRPCServer
	clients map[string]*GRPCClient
	mu      sync.RWMutex
}

// NewGRPCServerManager creates a new server and client manager
func NewGRPCServerManager() *GRPCServerManager {
	return &GRPCServerManager{
		servers: make(map[string]*GRPCServer),
		clients: make(map[string]*GRPCClient),
	}
}

// ============== SERVER MANAGEMENT ==============

// AddServer adds a new gRPC server to the manager
func (m *GRPCServerManager) AddServer(name string, server *GRPCServer) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.servers[name]; exists {
		return fmt.Errorf("server with name %s already exists", name)
	}
	m.servers[name] = server
	return nil
}

// StartServer starts a specific server
func (m *GRPCServerManager) StartServer(name string) error {
	m.mu.RLock()
	server, exists := m.servers[name]
	m.mu.RUnlock()

	if !exists {
		return fmt.Errorf("server with name %s not found", name)
	}

	go func() {
		if err := server.Start(); err != nil {
			log.Printf("Error starting server %s: %v", name, err)
		}
	}()

	log.Printf("Started gRPC server: %s", name)
	return nil
}

// StartAllServers starts all registered servers concurrently
// and waits until all servers are started before returning.
// If any server fails to start, it will panic immediately.
func (m *GRPCServerManager) StartAllServers() {
	m.mu.RLock()
	servers := make(map[string]*GRPCServer)
	for name, server := range m.servers {
		servers[name] = server
	}
	m.mu.RUnlock()

	var wg sync.WaitGroup
	serverErr := make(chan string, len(servers)) // buffered channel

	// Start servers concurrently
	for name, server := range servers {
		wg.Add(1)
		go func(serverName string, srv *GRPCServer) {
			defer wg.Done()
			log.Printf("🚀 Starting gRPC server: %s", serverName)
			if err := srv.Start(); err != nil {
				serverErr <- serverName
			}
		}(name, server)
	}

	// Wait for all servers to finish starting
	wg.Wait()
	close(serverErr)

	// Check if there was any error
	if len(serverErr) > 0 {
		for name := range serverErr {
			m.StopServer(name)
		}
		panic("start Grpc failed")
	}
}

// StopServer stops a specific server
func (m *GRPCServerManager) StopServer(name string) error {
	m.mu.RLock()
	server, exists := m.servers[name]
	m.mu.RUnlock()

	if !exists {
		return fmt.Errorf("server with name %s not found", name)
	}

	server.Stop()
	log.Printf("Stopped gRPC server: %s", name)
	return nil
}

// StopAllServers stops all servers
func (m *GRPCServerManager) StopAllServers() {
	m.mu.RLock()
	servers := make(map[string]*GRPCServer)
	for name, server := range m.servers {
		servers[name] = server
	}
	m.mu.RUnlock()

	for name, server := range servers {
		server.Stop()
		log.Printf("Stopped gRPC server: %s", name)
	}

	log.Println("All gRPC servers stopped")
}

// ============== CLIENT MANAGEMENT ==============

// AddClient adds a new gRPC client to the manager
func (m *GRPCServerManager) AddClient(name, addr string) (*GRPCClient, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if client, exists := m.clients[name]; exists {
		return client, nil
	}

	client, err := NewGRPCClient(addr)
	if err != nil {
		return nil, err
	}

	m.clients[name] = client
	log.Printf("Added gRPC client: %s -> %s", name, addr)
	return client, nil
}

// AddClientWithConfig adds a client with custom configuration
func (m *GRPCServerManager) AddClientWithConfig(name string, config *ClientConfig) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.clients[name]; exists {
		return fmt.Errorf("client with name %s already exists", name)
	}

	client, err := NewGRPCClientWithConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create client %s: %w", name, err)
	}

	m.clients[name] = client
	log.Printf("Added gRPC client: %s -> %s", name, config.Addr)
	return nil
}

// GetClient retrieves a client by name
func (m *GRPCServerManager) GetClient(name string) (*GRPCClient, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	client, exists := m.clients[name]
	if !exists {
		return nil, fmt.Errorf("client with name %s not found", name)
	}

	return client, nil
}

// CloseAllClients closes all clients
func (m *GRPCServerManager) CloseAllClients() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for name, client := range m.clients {
		if err := client.Close(); err != nil {
			log.Printf("Error closing client %s: %v", name, err)
		}
	}

	m.clients = make(map[string]*GRPCClient)
	log.Println("All gRPC clients closed")
}

// ============== LIFECYCLE MANAGEMENT ==============

// Shutdown gracefully shuts down all servers and clients
func (m *GRPCServerManager) Shutdown(ctx context.Context) error {
	log.Println("Starting graceful shutdown...")

	// Stop all servers
	m.StopAllServers()

	// Close all clients
	m.CloseAllClients()

	// Wait for shutdown or timeout
	select {
	case <-ctx.Done():
		log.Println("Shutdown timeout reached")
		return ctx.Err()
	case <-time.After(100 * time.Millisecond):
		log.Println("Graceful shutdown completed")
		return nil
	}
}
