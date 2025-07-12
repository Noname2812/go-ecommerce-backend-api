package commonenum

// AuthenticationState indicates the authentication state of a user.
type AuthenticationState int16

const (
	// NotAuthenticated indicates the user is not authenticated.
	NotAuthenticated AuthenticationState = 0

	// Pending indicates the user is pending authentication.
	Pending AuthenticationState = 1

	// Authenticated indicates the user is authenticated.
	Authenticated AuthenticationState = 2

	// Failed indicates the user authentication has failed.
	Failed AuthenticationState = 3
)

// IsValid checks if the AuthenticationState value is valid.
func (a AuthenticationState) IsValid() bool {
	return a == NotAuthenticated || a == Pending || a == Authenticated || a == Failed
}

// String converts the AuthenticationState value to a string.
func (a AuthenticationState) String() string {
	switch a {
	case NotAuthenticated:
		return "Not Authenticated"
	case Pending:
		return "Pending"
	case Authenticated:
		return "Authenticated"
	default:
		return "Failed"
	}
}
