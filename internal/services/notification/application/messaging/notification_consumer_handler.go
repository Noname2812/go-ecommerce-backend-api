package notificationmessaging

import "context"

type NotificationConsumer interface {
	Subscribe() error // Subscribe all topic of service
	// handler
	HandleOtpVerifyUserRegisterCreatedEvent(ctx context.Context, key, value []byte) error
}
