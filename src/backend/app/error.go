package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/go-chi/chi/v5/middleware"
)

type CustomErrors []*CustomError

func (e CustomErrors) Error() string {
	var sb strings.Builder
	for _, x := range e {
		sb.WriteString(x.Error())
	}
	return sb.String()
}

func NewCustomError(ctx context.Context, err error, errorCode CustomErrorCode, detail *CustomErrorDetail) *CustomError {
	ce := &CustomError{err: errors.WithStack(err), errorCode: errorCode, detail: detail}
	ce.setTraceID(ctx)
	return ce
}

func NewAuthenticationError(ctx context.Context, field string, val any) *CustomError {
	ce := &CustomError{errorCode: AuthenticationError, detail: NewCustomErrorDetail(field, val)}
	ce.setTraceID(ctx)
	return ce
}

func NewAuthorizationError(ctx context.Context, field string, val any) *CustomError {
	ce := &CustomError{errorCode: AuthorizationError, detail: NewCustomErrorDetail(field, val)}
	ce.setTraceID(ctx)
	return ce
}

func NewValidationError(ctx context.Context, err error, field string, val any) *CustomError {
	ce := &CustomError{err: errors.WithStack(err), errorCode: ValidationError, detail: NewCustomErrorDetail(field, val)}
	ce.setTraceID(ctx)
	return ce
}

func NewAlreadyExistsError(ctx context.Context, field string, val any) *CustomError {
	ce := &CustomError{errorCode: AlreadyExistsError, detail: NewCustomErrorDetail(field, val)}
	ce.setTraceID(ctx)
	return ce
}

func NewUnexpectedError(ctx context.Context, err error) *CustomError {
	ce := &CustomError{err: errors.WithStack(err), errorCode: UnexpectedError}
	ce.setTraceID(ctx)
	return ce
}

func NewUnexpectedErrorWithDetail(ctx context.Context, err error, field string, val any) *CustomError {
	ce := &CustomError{err: errors.WithStack(err), errorCode: UnexpectedError, detail: NewCustomErrorDetail(field, val)}
	ce.setTraceID(ctx)
	return ce
}

type CustomError struct {
	err       error
	traceID   string
	errorCode CustomErrorCode
	detail    *CustomErrorDetail
}

func (e *CustomError) setTraceID(ctx context.Context) {
	traceID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if ok {
		e.traceID = traceID
	}
}

func (e *CustomError) Error() string {
	var sb strings.Builder
	if e.err != nil {
		sb.WriteString(fmt.Sprintf("[err:%s]", e.err.Error()))
	}
	sb.WriteString(fmt.Sprintf("[error_code:%v]", e.errorCode))
	if e.traceID != "" {
		sb.WriteString(fmt.Sprintf("[trace_id:%s]", e.traceID))
	}
	if e.detail != nil {
		sb.WriteString(fmt.Sprintf("[field:%s][value:%v]", e.detail.Field, e.detail.Value))
	}
	return sb.String()
}

func (e *CustomError) GetCause() error {
	return errors.UnwrapAll(e.err)
}

func (e *CustomError) GetErrorCode() CustomErrorCode {
	return e.errorCode
}

func (e *CustomError) GetErrorDetail() *CustomErrorDetail {
	return e.detail
}

type CustomErrorCode string

func (c CustomErrorCode) ToString() string {
	return string(c)
}

const (
	// AuthenticationError is 認証エラー
	AuthenticationError CustomErrorCode = "AUTHENTICATION_ERROR"
	// AuthorizationError is 認可エラー
	AuthorizationError CustomErrorCode = "AUTHORIZATION_ERROR"
	// ValidationError is バリデーションエラー
	ValidationError CustomErrorCode = "VALIDATION_ERROR"
	// AlreadyExistsError is 存在チェックエラー
	AlreadyExistsError CustomErrorCode = "ALREADY_EXISTS_ERROR"

	// UnexpectedError is その他の予期せぬエラー
	UnexpectedError CustomErrorCode = "UNEXPECTED_ERROR"
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

func NewError(msg string) error {
	return errors.New(msg)
}

func WrapError(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func WrapErrorf(err error, format string, args ...any) error {
	return errors.Wrapf(err, format, args...)
}

func WithStackError(err error) error {
	return errors.WithStack(err)
}

func AsError(err error, target interface{}) bool {
	return errors.As(err, target)
}
