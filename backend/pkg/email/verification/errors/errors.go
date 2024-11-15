package errors

import "github.com/pkg/errors"

type EmailVerificationError struct {
	err error
}

func (e EmailVerificationError) Error() string {
	return e.err.Error()
}

var (
	ErrInvalidFormat = EmailVerificationError{errors.New("invalid emailVerification format")}
	ErrInvalidDomain = EmailVerificationError{errors.New("invalid emailVerification domain")}
	ErrEmailNotExist = EmailVerificationError{errors.New("email does not exist")}
)
