package mustache

// ErrorType is an integer type defining type of error
type ErrorType int

const (
	// ErrorTypeNone indicates a non-error.
	ErrorTypeNone ErrorType = iota
	// ErrorTypeUnknown indicates an error of unknown type.
	ErrorTypeUnknown
	// ErrorTypeMissingParams indicates that template parameters are missing.
	ErrorTypeMissingParams
)

type e struct {
	error
	errorType ErrorType
}

func WrapError(err error, errorType ErrorType) error {
	return e{
		err,
		errorType,
	}
}

// ErrorMatchesType returns true if provided error is of given type and false otherwise.
func ErrorMatchesType(err error, errorType ErrorType) bool {
	if err == nil {
		return errorType == ErrorTypeNone
	}
	e, ok := err.(e)
	if !ok {
		return errorType == ErrorTypeUnknown
	}
	return e.errorType == errorType
}
