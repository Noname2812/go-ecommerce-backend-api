package authmessaging

type AuthConsumer interface {
	Subscribe() error // Subscribe all topic of service
}
