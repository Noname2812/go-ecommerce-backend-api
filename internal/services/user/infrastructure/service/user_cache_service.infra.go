package userserviceimpl

import (
	userservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/service"
	"github.com/redis/go-redis/v9"
)

type userCacheService struct {
	rdb *redis.Client
}

// NewUserCacheService creates a new instance of UserCacheService.
func NewUserCacheService(redisClient *redis.Client) userservice.UserCacheService {
	return &userCacheService{
		rdb: redisClient,
	}
}
