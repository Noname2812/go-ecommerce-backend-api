package commonenum

// Gender indicates the gender of a user.
type Gender int16

const (
	// Secret indicates an undisclosed or unknown gender.
	Secret Gender = 0

	// Male indicates male gender.
	Male Gender = 1

	// Female indicates female gender.
	Female Gender = 2
)

// IsValid checks if the Gender value is valid.
func (g Gender) IsValid() bool {
	return g == Secret || g == Male || g == Female
}

// String converts the Gender value to a string.
func (g Gender) String() string {
	switch g {
	case Secret:
		return "Secret"
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Male"
	}
}
