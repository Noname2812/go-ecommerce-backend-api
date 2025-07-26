//go:build wireinject

package transportationwire

import (
	"database/sql"

	cacheservice "github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/cache"
	transportationqueryhandler "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/handler"
	transportationrepositoryimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/infrastructure/persistence"
	transportationserviceimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/infrastructure/service"
	"github.com/dgraph-io/ristretto"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var TransportationRepositorySet = wire.NewSet(
	transportationrepositoryimpl.NewTripRepository,
)

func InitTransportationQueryHandler(db *sql.DB, rdb *redis.Client, localCache *ristretto.Cache, logger *zap.Logger) transportationqueryhandler.TransportationQueryHandler {
	wire.Build(
		TransportationRepositorySet,
		cacheservice.NewRedisCache,
		cacheservice.NewLocalCache,
		transportationserviceimpl.NewTransportationQueryService,
		transportationqueryhandler.NewTransportationQueryHandler,
	)
	return nil
}
