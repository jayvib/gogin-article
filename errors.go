package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// This error handling pattern was taken from:
// https://hackernoon.com/golang-handling-errors-gracefully-8e27f1db729f

type ErrorType uint

const (
	NoType = ErrorType(iota)
	ItemNotFound
)

type ContextError struct {
	Field   string
	Message string
}

type customError struct {
	errorType     ErrorType
	originalError error
	contextInfo   ContextError
}

func (e customError) Error() string {
	return e.originalError.Error()
}

func (t ErrorType) New(msg string) error {
	return customError{
		errorType:     t,
		originalError: errors.New(msg),
	}
}

func (t ErrorType) Newf(msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return customError{
		errorType:     t,
		originalError: errors.New(msg),
	}
}

func (t ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, args...)
	return customError{
		errorType:     t,
		originalError: newErr,
	}
}

func New(msg string) error {
	return customError{
		errorType:     NoType,
		originalError: errors.New(msg),
	}
}

func Newf(msg string, args ...interface{}) error {
	return customError{
		errorType:     NoType,
		originalError: errors.New(fmt.Sprintf(msg, args...)),
	}
}

func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedErr := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: wrappedErr,
			contextInfo:   customErr.contextInfo,
		}
	}
	return customError{
		errorType:     NoType,
		originalError: wrappedErr,
	}
}

func AddContextError(err error, field, message string) error {
	context := ContextError{
		Field:   field,
		Message: message,
	}
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: customErr.originalError,
			contextInfo:   context,
		}
	}
	return customError{
		errorType:     NoType,
		originalError: err,
		contextInfo:   context,
	}
}

func GetErrorContext(err error) map[string]string {
	emptyContext := ContextError{}
	if customErr, ok := err.(customError); ok || customErr.contextInfo != emptyContext {
		return map[string]string{
			"field":   customErr.contextInfo.Field,
			"message": customErr.contextInfo.Message,
		}
	}
	return nil
}
