package commonenum

// UserState indicates the state of a user account
type UserState int16

const (
	Locked UserState = 0 // Locked: Account is locked or disabled

	Activated UserState = 1 // Activated: Account is active and usable

	NotActivated UserState = 2 // NotActivated: Account has not been activated yet
)
