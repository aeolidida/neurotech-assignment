package errs

import (
	"errors"
	"fmt"
)

var (
	ErrNoPatient = errors.New("no such patient")
)

var (
	ErrInvalidGender = errors.New("invalid gender")
	ErrInvalidDate   = errors.New("invalid date")
)

func WrapError(op, message string, err error) error {
	return fmt.Errorf("%s %s: %w", op, message, err)
}
