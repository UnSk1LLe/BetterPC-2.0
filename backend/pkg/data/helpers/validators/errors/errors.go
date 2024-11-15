package errors

import (
	"github.com/pkg/errors"
)

type ValidatorError struct {
	err error
}

func (v ValidatorError) Error() string {
	return v.err.Error()
}

var (
	ErrEmptyStruct  = ValidatorError{errors.New("all item fields are empty")}
	ErrInvalidInput = ValidatorError{errors.New("input must be a pointer or a struct")}
)
