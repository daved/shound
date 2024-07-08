package ccmd

import (
	"errors"
	"fmt"
)

var ErrHelpFlag = errors.New("help requested")

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

type NotInstalledError struct {
	theme string
}

func NewNotInstalledError(theme string) *NotInstalledError {
	return &NotInstalledError{theme}
}

func (e *NotInstalledError) Error() string {
	return fmt.Sprintf("%q is not a valid installed theme", e.theme)
}

type AlreadyInstalledError struct {
	theme string
}

func NewAlreadyInstalledError(theme string) *AlreadyInstalledError {
	return &AlreadyInstalledError{theme}
}

func (e *AlreadyInstalledError) Error() string {
	return fmt.Sprintf("the theme %q is already installed", e.theme)
}
