package app

import "github.com/pkg/errors"

type CustomError interface {
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

type customError struct {
	err error
}

func (e *customError) Error() string {
	return e.err.Error()
}
