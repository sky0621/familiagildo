package controller

import (
	"context"
	"errors"
	"fmt"
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
	details []CustomErrorDetail
}

type CustomErrorDetail struct {
	Field string
	Value any
}

func (e *CustomError) AddGraphQLError(ctx context.Context, msg string) {
	extensions := map[string]interface{}{
		"status_code": e.httpStatusCode,
		"error_code":  e.appErrorCode,
	}
	for i, d := range e.details {
		extensions[fmt.Sprintf("field_%d", i+1)] = d.Field
		extensions[fmt.Sprintf("value_%d", i+1)] = d.Value
	}
	graphql.AddError(ctx, &gqlerror.Error{
		Message:    msg,
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
func ValidationError(details []CustomErrorDetail, opts ...CustomErrorOption) *CustomError {
	var options []CustomErrorOption
	for _, d := range details {
		options = append(options, WithCustomErrorDetail(d))
	}
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

func WithCustomErrorDetail(v CustomErrorDetail) CustomErrorOption {
	return func(a *CustomError) {
		a.details = append(a.details, v)
	}
}

func AddGraphQLError(ctx context.Context, err error) {
	var anErr *app.AuthenticationError
	if errors.As(err, &anErr) {
		AuthenticationError().AddGraphQLError(ctx, "認証に失敗しました。") // FIXME: i18n
		return
	}

	var azErr *app.AuthorizationError
	if errors.As(err, &azErr) {
		AuthorizationError(WithCustomErrorDetail(CustomErrorDetail{
			Field: "userID", Value: azErr.GetUserID(),
		})).AddGraphQLError(ctx, "認可に失敗しました。") // FIXME: i18n
		return
	}

	var vnErr *app.ValidationError
	if errors.As(err, &vnErr) {
		var cErrs []CustomErrorDetail
		for _, d := range vnErr.GetDetails() {
			cErrs = append(cErrs, CustomErrorDetail{Field: d.GetField(), Value: d.GetValue()})
		}
		ValidationError(cErrs, WithCustomErrorDetail(CustomErrorDetail{
			Field: "userID", Value: vnErr.GetUserID(),
		})).AddGraphQLError(ctx, "バリデーションに失敗しました。") // FIXME: i18n
		return
	}

	var uErr *app.UnexpectedError
	if errors.As(err, &uErr) {
		InternalServerError(WithCustomErrorDetail(CustomErrorDetail{
			Field: "userID", Value: uErr.GetUserID(),
		})).AddGraphQLError(ctx, "予期せぬエラーが発生しました。") // FIXME: i18n
		return
	}

	cErr := &CustomError{}
	cErr.AddGraphQLError(ctx, err.Error()) // FIXME: i18n

}
