package app

import "github.com/cockroachdb/errors"

const (
	// Unknown is ユーザーID不明
	// FIXME:
	Unknown = -1
)

func NewAuthenticationError(err error, msg string) *AuthenticationError {
	return &AuthenticationError{err: errors.Wrap(err, msg)}
}

type AuthenticationError struct {
	err error
}

func (e *AuthenticationError) Error() string {
	return e.err.Error()
}

func NewCustomError(err error, msg string) *CustomError {
	return &CustomError{err: errors.Wrap(err, msg)}
}

type CustomError struct {
	err        error
	supplement CustomErrorSupplement
}

type CustomErrorSupplement interface {
}

// AuthenticationErrorSupplement is 認証エラー補足情報
type AuthenticationErrorSupplement struct {
	userID int
}

func NewAuthenticationErrorSupplement(userID int) CustomErrorSupplement {
	return &AuthenticationErrorSupplement{userID: userID}
}

// AuthorizationErrorSupplement is 認可エラー補足情報
type AuthorizationErrorSupplement struct {
	userID int
	funcID int
}

func NewAuthorizationErrorSupplement(userID, funcID int) CustomErrorSupplement {
	return &AuthorizationErrorSupplement{userID: userID, funcID: funcID}
}

// ValidationErrorSupplement is バリデーションエラー補足情報
type ValidationErrorSupplement struct {
	userID  int
	details []ValidationErrorDetail
}

type ValidationErrorDetail struct {
	field string
	value any
}

func NewValidationErrorSupplement(userID int, details []ValidationErrorDetail) CustomErrorSupplement {
	return &ValidationErrorSupplement{userID: userID, details: details}
}

// UnexpectedErrorSupplement is 予期せぬエラー補足情報
type UnexpectedErrorSupplement struct {
	userID int
}

func NewUnexpectedErrorSupplement(userID int) CustomErrorSupplement {
	return &UnexpectedErrorSupplement{userID: userID}
}
