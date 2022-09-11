package handler

import "net/http"

var (
	// ErrUnhealthy is 500 error, which means that something went wrong.
	ErrUnhealthy = NewError(http.StatusInternalServerError, "system.unhealthy")

	ErrInvalidBody  = NewError(http.StatusBadRequest, "request.invalid_body")
	ErrUnauthorized = NewError(http.StatusUnauthorized, "user.unauthorized")
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents HTTP error.
type StatusError struct {
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
}

// Error allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Messages[0]
}

// Status returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// NewError creates new error instance.
func NewError(code int, messages ...string) Error {
	return StatusError{
		Code:     code,
		Messages: messages,
	}
}
