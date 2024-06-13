package ccmd

import (
	"errors"
	"fmt"
)

var ErrHelpFlag = errors.New("help requested")

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
