package commonenum

// EventStatus indicates the status of an event
type EventStatus uint8

const (
	EVENTPENDING EventStatus = 1 // PENDING: Event is pending

	EVENTPUBLISHED EventStatus = 2 // PUBLISHED: Event is published

	EVENTFAILED EventStatus = 3 // FAILED: Event has failed
)
