package errors

import "github.com/pkg/errors"

type UserError struct {
	err error
}

func (u UserError) Error() string {
	return u.err.Error()
}

var (
	ErrUserAlreadyExists   = UserError{errors.New("user already exists")}
	ErrUserNotFound        = UserError{errors.New("user not found")}
	ErrUserAlreadyVerified = UserError{errors.New("user already verified")}
	ErrUserNotVerified     = UserError{errors.New("user is not verified")}
	ErrTokenExpired        = UserError{errors.New("token expired")}
	ErrInvalidUserRole     = UserError{errors.New("invalid user role")}
	ErrUserNotModified     = UserError{errors.New("user not modified, no new changes for user")}
)
