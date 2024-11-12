package cmd

import (
	"fmt"
)

type UsageError struct {
	err error
}

func NewUsageError(err error) *UsageError {
	return &UsageError{
		err: err,
	}
}

func (e *UsageError) Error() string {
	return fmt.Sprintf("usage: %v", e.err)
}

func (e *UsageError) Unwrap() error {
	return e.err
}
