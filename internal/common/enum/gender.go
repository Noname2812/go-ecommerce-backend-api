package commonenum

// Gender indicates the gender of a user.
type Gender int16

const (
	Secret Gender = 0 // Secret: An undisclosed or unknown gender.

	Male Gender = 1 // Male: Male gender.

	Female Gender = 2 // Female: Female gender.
)
