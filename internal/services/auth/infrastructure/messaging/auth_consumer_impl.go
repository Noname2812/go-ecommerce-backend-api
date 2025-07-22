package authmessagingimpl

import authmessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/messaging"

type authConsumer struct {
}

// Subscribe implements authmessaging.AuthConsumerHandler.
func (a *authConsumer) Subscribe() error {
	panic("unimplemented")
}

func NewAuthConsumer() authmessaging.AuthConsumer {
	return &authConsumer{}
}
