package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

//Reference article:
//https://hackernoon.com/golang-handling-errors-gracefully-8e27f1db729f

type CustomError struct {
	errType      ErrorType
	wrappedError error
	context      errorContext
	opt          *option
}

type option struct {
	stacktrace bool
}

type Option func(*option)

func WithStackTrace() Option {
	return func(opt *option) {
		opt.stacktrace = true
	}
}

func (err CustomError) Error() string {
	if err.opt != nil && err.opt.stacktrace {
		return err.Stacktrace()
	}
	return err.wrappedError.Error()
}

func (err CustomError) Stacktrace() string {
	return fmt.Sprintf("%+v\n", err.wrappedError)
}

// New creates a no type error
func New(msg string, opts ...Option) error {
	o := option{}
	for _, opt := range opts {
		opt(&o)
	}

	return CustomError{errType: Unknown, wrappedError: errors.New(msg), opt: &o}
}

// Newf creates a no type error with formatted message
func Newf(msg string, args ...interface{}) error {
	return CustomError{errType: Unknown, wrappedError: errors.New(fmt.Sprintf(msg, args...))}
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

	return CustomError{errType: Unknown, wrappedError: wrappedError}
}

// Get Stacktrace of error
func Stack(err error) string {
	if customErr, ok := err.(CustomError); ok {
		return fmt.Sprintf("%+v\n", customErr.wrappedError)
	}
	return fmt.Sprintf("%+v\n", errors.WithStack(err))
}
