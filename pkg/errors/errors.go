package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

//Reference article:
//https://hackernoon.com/golang-handling-errors-gracefully-8e27f1db729f

// CustomError struct
type CustomError struct {
	errType      ErrorType
	wrappedError error
	context      errorContext
}

// Error return error message
func (err CustomError) Error() string {
	return err.wrappedError.Error()
}

// Stacktrace return error stacktrace message
func (err CustomError) Stacktrace() string {
	return fmt.Sprintf("%+v\n", err.wrappedError)
}

// New creates a no type error
func New(msg string) error {
	return CustomError{errType: Error, wrappedError: errors.New(msg)}
}

// Newf creates a no type error with formatted message
func Newf(msg string, args ...interface{}) error {
	return CustomError{errType: Error, wrappedError: errors.New(fmt.Sprintf(msg, args...))}
}

// Wrap wrans an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Cause gives the original error
func Cause(err error) error {
	return errors.Cause(err)
}

// Wrapf wraps an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(CustomError); ok {
		return CustomError{
			errType:      customErr.errType,
			wrappedError: wrappedError,
			context:      customErr.context,
		}
	}

	return CustomError{errType: Error, wrappedError: wrappedError}
}

// Stack get stacktrace of error
func Stack(err error) string {
	if customErr, ok := err.(CustomError); ok {
		return fmt.Sprintf("%+v\n", customErr.wrappedError)
	}
	return fmt.Sprintf("%+v\n", errors.WithStack(err))
}
