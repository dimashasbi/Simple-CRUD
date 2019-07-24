package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	// NoType for something
	NoType = ErrorType(iota)
	// BadRequest for
	BadRequest
	// NotFound for
	NotFound
	//add any type you want
)

// ErrorType ..
type ErrorType uint

type customError struct {
	errorType     ErrorType
	originalError error
	contextInfo   map[string]string
}

// Error returns the mssage of a customError
func (er customError) Error() string {
	return er.originalError.Error()
}

// New creates a new customError
func (er ErrorType) New(msg string) error {
	return customError{errorType: er, originalError: errors.New(msg)}
}

// Newf missing ... in args forwarded to printf-like functiongo-vet creates a new customError with formatted message
func (er ErrorType) Newf(msg string, args ...interface{}) error {
	err := fmt.Errorf(msg, er)

	return customError{errorType: er, originalError: err}
}

// Wrap creates a new wrapped error
func (er ErrorType) Wrap(err error, msg string) error {
	return er.Wrapf(err, msg)
}

// Wrapf creates a new wrapped error with formatted message
func (er ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, er)

	return customError{errorType: er, originalError: newErr}
}
