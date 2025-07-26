package usermessaging

import "context"

// UserConsumer is the interface for the user consumer
type UserConsumer interface {
	// Subscribe all topic of service
	Subscribe() error
	// handle user base inserted
	HandleUserBaseInserted(ctx context.Context, key, value []byte) error
}
