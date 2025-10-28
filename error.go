// Package gn provides user-friendly messaging and error handling utilities
// for Go applications with colorized output and formatted messages.
package gn

import (
	"errors"
	"fmt"
)

// PrintErrorMessage attempts to print an error message if the error implements
// the Error interface. It uses errors.As to check if the error can be converted
// to the custom Error type, and if so, prints it with the appropriate formatting.
func PrintErrorMessage(err error) {
	var target *Error
	if errors.As(err, &target) {
		target.print()
	}
	Message("\n<err>Error:</err> " + err.Error())
}

// ErrorCode represents a numeric code for categorizing errors.
type ErrorCode int

// Error is a custom error type that extends the standard error interface
// with additional fields for error codes, custom messages, and formatting variables.
// It can be used to create rich, user-friendly error messages.
type Error struct {
	Code ErrorCode // Numeric code identifying the type of error
	Err  error     // The underlying error
	Msg  string    // User-friendly message template (can include format specifiers)
	Vars []any     // Variables to be formatted into the message template
}

// Error implements the error interface, returning the underlying error's message
// if available, or the custom message otherwise.
func (e *Error) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	if len(e.Vars) > 0 {
		return fmt.Sprintf(e.Msg, e.Vars...)
	}
	return e.Msg
}

// print formats and displays the error message with the error icon and styling.
// This is an internal method called by PrintErrorMessage.
func (e *Error) print() {
	print(errorMsgType, e.Msg, e.Vars)
}
