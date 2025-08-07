package initialize

import (
	"sync"
)

var (
	serviceRegistry = make(map[string]interface{})
	registryMutex   sync.RWMutex
)

func RegisterService(name string, service interface{}) {
	registryMutex.Lock()
	defer registryMutex.Unlock()
	serviceRegistry[name] = service
}

func GetService(name string) interface{} {
	registryMutex.RLock()
	defer registryMutex.RUnlock()
	return serviceRegistry[name]
}

func CleanupServices() {
	registryMutex.Lock()
	defer registryMutex.Unlock()

	for _, service := range serviceRegistry {
		if closer, ok := service.(interface{ Close() error }); ok {
			closer.Close()
		}
	}
	serviceRegistry = make(map[string]interface{})
}
