package app

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/go-chi/chi/v5/middleware"
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

func NewCustomError(ctx context.Context, err error, errorCode CustomErrorCode, detail *CustomErrorDetail) CustomError {
	traceID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if ok {
		return CustomError{err: errors.WithStack(err), traceID: traceID, errorCode: errorCode, detail: detail}
	}
	return CustomError{err: errors.WithStack(err), errorCode: errorCode, detail: detail}
}

func NewValidationError(ctx context.Context, err error, field string, val any) CustomError {
	traceID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if ok {
		return CustomError{err: errors.WithStack(err), traceID: traceID, errorCode: ValidationError, detail: NewCustomErrorDetail(field, val)}
	}
	return CustomError{err: errors.WithStack(err), errorCode: ValidationError, detail: NewCustomErrorDetail(field, val)}
}

func NewAlreadyExistsError(ctx context.Context, field string, val any) CustomError {
	traceID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if ok {
		return CustomError{traceID: traceID, errorCode: AlreadyExistsError, detail: NewCustomErrorDetail(field, val)}
	}
	return CustomError{errorCode: AlreadyExistsError, detail: NewCustomErrorDetail(field, val)}
}

func NewUnexpectedError(ctx context.Context, err error) CustomError {
	traceID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if ok {
		return CustomError{err: errors.WithStack(err), traceID: traceID, errorCode: UnexpectedError}
	}
	return CustomError{err: errors.WithStack(err), errorCode: UnexpectedError}
}

func NewUnexpectedErrorWithDetail(ctx context.Context, err error, field string, val any) CustomError {
	return CustomError{err: errors.WithStack(err), errorCode: UnexpectedError, detail: NewCustomErrorDetail(field, val)}
}

type CustomError struct {
	err       error
	traceID   string
	errorCode CustomErrorCode
	detail    *CustomErrorDetail
}

func (e CustomError) Error() string {
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

func (e CustomError) GetCause() error {
	return errors.UnwrapAll(e.err)
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

func WrapError(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func WrapErrorf(err error, format string, args ...any) error {
	return errors.Wrapf(err, format, args...)
}
