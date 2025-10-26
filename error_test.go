package gn

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *Error
		expected string
	}{
		{
			name: "with underlying error",
			err: &Error{
				Code: 1,
				Err:  errors.New("underlying error"),
				Msg:  "custom message",
			},
			expected: "underlying error",
		},
		{
			name: "with message and vars",
			err: &Error{
				Code: 2,
				Msg:  "error: %s, code: %d",
				Vars: []any{"test", 42},
			},
			expected: "error: test, code: 42",
		},
		{
			name: "with message only",
			err: &Error{
				Code: 3,
				Msg:  "simple error message",
			},
			expected: "simple error message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Error()
			if result != tt.expected {
				t.Errorf("Error() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestPrintErrorMessage(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "with Error type",
			err: &Error{
				Code: 1,
				Msg:  "test error message",
			},
		},
		{
			name: "with standard error",
			err:  errors.New("standard error"),
		},
		{
			name: "with wrapped Error",
			err: fmt.Errorf("wrapped: %w", &Error{
				Code: 2,
				Msg:  "wrapped error",
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test primarily checks that PrintErrorMessage doesn't panic
			// Actual output testing would require capturing stdout
			PrintErrorMessage(tt.err)
		})
	}
}

func TestError_ImplementsErrorInterface(t *testing.T) {
	var _ error = &Error{}
	var _ error = (*Error)(nil)
}

func TestErrorCode(t *testing.T) {
	var code ErrorCode = 42
	if code != 42 {
		t.Errorf("ErrorCode = %d, want 42", code)
	}
}

func TestErrorsAs(t *testing.T) {
	// Test that errors.As works correctly with our Error type
	originalErr := &Error{
		Code: 100,
		Msg:  "original error",
	}

	wrappedErr := fmt.Errorf("wrapped: %w", originalErr)

	var target *Error
	if !errors.As(wrappedErr, &target) {
		t.Fatal("errors.As failed to unwrap Error type")
	}

	if target.Code != 100 {
		t.Errorf("target.Code = %d, want 100", target.Code)
	}

	if target.Msg != "original error" {
		t.Errorf("target.Msg = %q, want %q", target.Msg, "original error")
	}
}
