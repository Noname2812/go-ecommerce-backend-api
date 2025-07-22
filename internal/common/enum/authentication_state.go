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
