package schema

import (
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/schema/codes"
	"github.com/aidenwallis/go-utils/utils"
)

var (
	// InvalidTokenError is returned when the session token is invalid
	InvalidTokenError = NewError(codes.InvalidToken, "Invalid token")

	// InvalidBodyError is returned when we cannot decode the request body
	InvalidBodyError = NewError(codes.InvalidBody, "Invalid body")

	// UnknownError is the error instance used as a fallback error
	UnknownError = NewError(codes.Unknown, "Something went wrong, please try again later.")
)

// Error defines the error body
type Error struct {
	Code     codes.ErrorCode `json:"code"`
	Message  string          `json:"message"`
	Metadata interface{}     `json:"metadata,omitempty"`
}

// NewError creates a new error
func NewError(code codes.ErrorCode, message string, metadata ...interface{}) *Error {
	return &Error{
		Code:     code,
		Message:  message,
		Metadata: utils.FirstOf(metadata),
	}
}
