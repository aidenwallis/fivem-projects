package codes

// ErrorCode defines all available error codes
type ErrorCode int

const (
	// Unknown is the code used for unknown errors.
	Unknown = ErrorCode(0)

	// InvalidBody is the code used when we cannot decode the request body
	InvalidBody = ErrorCode(1)

	// ValidationError is the code used when the request body fails validation
	ValidationError = ErrorCode(2)

	// InvalidToken is the code used when the token is invalid.
	InvalidToken = ErrorCode(3)
)
