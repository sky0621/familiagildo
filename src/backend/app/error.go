package app

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"strings"
)

type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	var sb strings.Builder
	for _, x := range e {
		sb.WriteString(x.Error())
	}
	return sb.String()
}

func NewValidationError(err error, detail *ValidationErrorDetail) ValidationError {
	return ValidationError{err: err, detail: detail}
}

type ValidationError struct {
	err    error
	detail *ValidationErrorDetail
}

type ValidationErrorDetail struct {
	Field string
	Value any
}

func (e *ValidationError) Error() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("[Field:%s, Value:%v]", e.detail.Field, e.detail.Value))
	return errors.WithDetailf(e.err, "%s", sb.String()).Error()
}

func (e *ValidationError) GetDetail() *ValidationErrorDetail {
	return e.detail
}

// NewUnexpectedError is 予期せぬエラーを生成
func NewUnexpectedError(err error, message string) *UnexpectedError {
	return &UnexpectedError{err: errors.WithStack(err), message: message}
}

type UnexpectedError struct {
	err     error
	message string
}

func (e *UnexpectedError) Error() string {
	return errors.WithDetailf(e.err, "%s", e.message).Error()
}

/*
 * ユーティリティ
 */

func WrapError(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func WrapErrorf(err error, format string, args ...any) error {
	return errors.Wrapf(err, format, args...)
}
