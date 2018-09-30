package interfaces

import "errors"

var (
	// ErrInternalServerError provides 500 error
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound provides 404 error
	ErrNotFound = errors.New("Your requested Item is not found")
)
