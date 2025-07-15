package usermessaging

import "context"

type UserConsumerHandler interface {
	Subscribe() error // Subscribe all topic of service
	// hanlder
	HandleUserBaseInserted(ctx context.Context, key, value []byte) error
}
