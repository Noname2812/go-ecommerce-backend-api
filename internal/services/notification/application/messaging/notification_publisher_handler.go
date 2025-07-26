package notificationmessaging

type NotificationPublisher interface {
	Register() // register all producers of service
	// handlers
}
