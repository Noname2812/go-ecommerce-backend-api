package errors

func MustCheck(err error, msg string) {
	if err != nil {
		panic(err)
	}
}

// not dependentcy looger
func Must(err error, component string) {
	if err != nil {
		panic(err)
	}
}
