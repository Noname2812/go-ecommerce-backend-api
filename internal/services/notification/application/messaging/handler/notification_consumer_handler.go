package notificationmessaginghandler

import (
	"context"
)

type NotificationConsumerHandler interface {
	HandleOtpVerifyUserRegisterCreatedEvent(ctx context.Context, key, value []byte) error
}
