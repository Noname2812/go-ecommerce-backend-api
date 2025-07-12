package commonenum

// UserState indicates the state of a user account.
type UserState int16

const (
	// Locked indicates the account is locked or disabled.
	Locked UserState = 0

	// Activated indicates the account is active and usable.
	Activated UserState = 1

	// NotActivated indicates the account has not been activated yet.
	NotActivated UserState = 2
)

// IsValid checks if the UserState value is valid.
func (u UserState) IsValid() bool {
	return u == Locked || u == Activated || u == NotActivated
}

// String converts the UserState value to a string.
func (u UserState) String() string {
	switch u {
	case Locked:
		return "Locked"
	case Activated:
		return "Activated"
	case NotActivated:
		return "Not Activated"
	default:
		return "Locked"
	}
}
