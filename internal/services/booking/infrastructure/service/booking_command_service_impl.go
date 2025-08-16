package bookingserviceimpl

import (
	"context"

	cacheservice "github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/cache"
	bookingcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/application/command/dto/request"
	bookingservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/application/service"
	bookingrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/domain/repository"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
)

type bookingCommandServiceImpl struct {
	logger                   *zap.Logger
	bookingCommandRepository bookingrepository.BookingRepository
	redisCacheService        cacheservice.RedisCache
	sfGroup                  singleflight.Group
}

// CreateBooking implements bookingservice.BookingCommandService.
func (b *bookingCommandServiceImpl) CreateBooking(ctx context.Context, body *bookingcommandrequest.CreateBookingRequest) (int, error, interface{}) {
	// bookingID := uuid.New().String()

	// 1. Call Transportation gRPC service to lock seats
	return 0, nil, nil
}

func NewBookingCommandServiceImpl(
	bookingCommandRepository bookingrepository.BookingRepository,
	redisCacheService cacheservice.RedisCache,

	logger *zap.Logger,
) bookingservice.BookingCommandService {
	return &bookingCommandServiceImpl{
		bookingCommandRepository: bookingCommandRepository,
		redisCacheService:        redisCacheService,
		sfGroup:                  singleflight.Group{},
		logger:                   logger,
	}
}
