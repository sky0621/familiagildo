package controller

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/sky0621/familiagildo/app"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
)

type CustomError struct {
	httpStatusCode int             // http.StatusCodeXXXXXXX を入れる
	appErrorCode   CustomErrorCode // サービス固有に定義したエラーコード

	/*
	 * 以下、全てのエラー表現に必須ではない要素（オプションとして設定可能）
	 */
	field string
	value string
}

func (e *CustomError) AddGraphQLError(ctx context.Context) {
	extensions := map[string]interface{}{
		"status_code": e.httpStatusCode,
		"error_code":  e.appErrorCode,
	}
	if e.field != "" {
		extensions["field"] = e.field
	}
	if e.value != "" {
		extensions["value"] = e.value
	}
	graphql.AddError(ctx, &gqlerror.Error{
		Message:    "",
		Extensions: extensions,
	})
}

func NewCustomError(httpStatusCode int, appErrorCode CustomErrorCode, opts ...CustomErrorOption) *CustomError {
	a := &CustomError{
		httpStatusCode: httpStatusCode,
		appErrorCode:   appErrorCode,
	}

	for _, o := range opts {
		o(a)
	}

	return a
}

// AuthenticationError is 認証エラー用
func AuthenticationError(opts ...CustomErrorOption) *CustomError {
	return NewCustomError(http.StatusUnauthorized, AuthenticationFailure, opts...)
}

// AuthorizationError is 認可エラー用
func AuthorizationError(opts ...CustomErrorOption) *CustomError {
	return NewCustomError(http.StatusForbidden, AuthorizationFailure, opts...)
}

// ValidationError is バリデーションエラー用
func ValidationError(field, value string, opts ...CustomErrorOption) *CustomError {
	options := []CustomErrorOption{WithField(field), WithValue(value)}
	for _, opt := range opts {
		options = append(options, opt)
	}
	return NewCustomError(http.StatusBadRequest, ValidationFailure, options...)
}

// InternalServerError is その他エラー用
func InternalServerError(opts ...CustomErrorOption) *CustomError {
	return NewCustomError(http.StatusInternalServerError, UnexpectedFailure, opts...)
}

type CustomErrorCode string

// MEMO: サービスの定義によっては意味のある文字列よりもコード体系を決めるのもあり。
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

type CustomErrorOption func(*CustomError)

func WithField(v string) CustomErrorOption {
	return func(a *CustomError) {
		a.field = v
	}
}

func WithValue(v string) CustomErrorOption {
	return func(a *CustomError) {
		a.value = v
	}
}

func AddGraphQLError(ctx context.Context, tag app.CustomErrorTag, opts ...CustomErrorOption) {
	switch tag {
	case app.AuthenticationFailure:
		AuthenticationError(opts...).AddGraphQLError(ctx)
	case app.AuthorizationFailure:
		AuthorizationError(opts...).AddGraphQLError(ctx)
	case app.ValidationFailure:
		ValidationError("", "", opts...).AddGraphQLError(ctx)
	default:
		InternalServerError(opts...).AddGraphQLError(ctx)
	}
}
