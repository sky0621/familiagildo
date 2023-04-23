package app

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"strings"
)

const (
	// Unknown is ユーザーID不明
	// FIXME:
	Unknown = -1
)

// NewAuthenticationError is 認証エラーを生成
func NewAuthenticationError(err error, userID int) *AuthenticationError {
	return &AuthenticationError{err: errors.WithStack(err), userID: userID}
}

type AuthenticationError struct {
	err    error
	userID int
}

func (e *AuthenticationError) Error() string {
	return errors.WithDetailf(e.err, "[UserID:%d]", e.userID).Error()
}

// NewAuthorizationError is 認可エラーを生成
func NewAuthorizationError(err error, userID, funcID int) *AuthorizationError {
	return &AuthorizationError{err: errors.WithStack(err), userID: userID, funcID: funcID}
}

type AuthorizationError struct {
	err    error
	userID int
	funcID int
}

func (e *AuthorizationError) Error() string {
	return errors.WithDetailf(e.err, "[UserID:%d][FuncID:%d]", e.userID, e.funcID).Error()
}

// NewValidationError is バリデーションエラーを生成
func NewValidationError(err error, userID int, details []ValidationErrorDetail) *ValidationError {
	return &ValidationError{err: errors.WithStack(err), userID: userID, details: details}
}

type ValidationError struct {
	err     error
	userID  int
	details []ValidationErrorDetail
}

type ValidationErrorDetail struct {
	field string
	value any
}

func (d *ValidationErrorDetail) GetField() string {
	return d.field
}

func (d *ValidationErrorDetail) GetValue() any {
	return d.value
}

func NewValidationErrorDetail(field string, value any) ValidationErrorDetail {
	return ValidationErrorDetail{field: field, value: value}
}

func (e *ValidationError) Error() string {
	sb := strings.Builder{}
	for _, d := range e.details {
		sb.WriteString(fmt.Sprintf("[field:%s, value:%v]", d.field, d.value))
	}
	return errors.WithDetailf(e.err, "[UserID:%d]%s", e.userID, sb.String()).Error()
}

func (e *ValidationError) GetDetails() []ValidationErrorDetail {
	return e.details
}

// NewUnexpectedError is 予期せぬエラーを生成
func NewUnexpectedError(err error, userID int, message string) *UnexpectedError {
	return &UnexpectedError{err: errors.WithStack(err), userID: userID, message: message}
}

type UnexpectedError struct {
	err     error
	userID  int
	message string
}

func (e *UnexpectedError) Error() string {
	return errors.WithDetailf(e.err, "%s [UserID:%d]", e.message, e.userID).Error()
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
