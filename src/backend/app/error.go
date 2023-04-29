package app

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"strings"
)

type CustomErrors []CustomError

func (e CustomErrors) Error() string {
	var sb strings.Builder
	for _, x := range e {
		sb.WriteString(x.Error())
	}
	return sb.String()
}

func NewCustomError(err error, errorCode CustomErrorCode, detail *CustomErrorDetail) CustomError {
	return CustomError{err: errors.WithStack(err), errorCode: errorCode, detail: detail}
}

type CustomError struct {
	err       error
	errorCode CustomErrorCode
	detail    *CustomErrorDetail
}

func (e CustomError) Error() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[err:%s]", e.err.Error()))
	sb.WriteString(fmt.Sprintf("[error_code:%v]", e.errorCode))
	sb.WriteString(fmt.Sprintf("[field:%s][value:%v]", e.detail.Field, e.detail.Value))
	return sb.String()
}

func (e CustomError) GetCause() error {
	return e.err
}

func (e CustomError) GetErrorCode() CustomErrorCode {
	return e.errorCode
}

func (e CustomError) GetErrorDetail() *CustomErrorDetail {
	return e.detail
}

type CustomErrorCode string

func (c CustomErrorCode) ToString() string {
	return string(c)
}

const (
	// AuthenticationFailure is 認証エラー
	AuthenticationFailure CustomErrorCode = "AUTHENTICATION_FAILURE"
	// AuthorizationFailure is 認可エラー
	AuthorizationFailure CustomErrorCode = "AUTHORIZATION_FAILURE"
	// ValidationFailure is バリデーションエラー
	ValidationFailure CustomErrorCode = "VALIDATION_FAILURE"

	// UnexpectedFailure is その他の予期せぬエラー
	UnexpectedFailure CustomErrorCode = "UNEXPECTED_FAILURE"
)

func NewCustomErrorDetail(field string, value any) *CustomErrorDetail {
	return &CustomErrorDetail{Field: field, Value: value}
}

type CustomErrorDetail struct {
	Field string
	Value any
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
