package app

import "github.com/pkg/errors"

type CustomError interface {
	Equals(tag CustomErrorTag) bool
	Error() string
}

func WrapError(err error) CustomError {
	return &customError{err: errors.WithStack(err)}
}

func WrapErrorWithMsg(err error, msg string) CustomError {
	return &customError{err: errors.Wrap(err, msg)}
}

func WrapErrorWithMsgf(err error, format string, args ...any) CustomError {
	return &customError{err: errors.Wrapf(err, format, args...)}
}

type CustomErrorTag int

const (
	// AuthenticationFailure is 認証エラー
	AuthenticationFailure CustomErrorTag = iota + 1
	// AuthorizationFailure is 認可エラー
	AuthorizationFailure
	// ValidationFailure is バリデーションエラー
	ValidationFailure

	// UnexpectedFailure is その他の予期せぬエラー
	UnexpectedFailure
)

type customError struct {
	tag CustomErrorTag
	err error
}

func (e *customError) Equals(tag CustomErrorTag) bool {
	if e == nil {
		return false
	}
	return e.tag == tag
}

func (e *customError) Error() string {
	return e.err.Error()
}
