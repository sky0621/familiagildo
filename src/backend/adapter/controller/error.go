package controller

import (
	"context"
	"errors"
	"github.com/sky0621/familiagildo/app"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func toMapFromCustomError(e app.CustomError) map[string]any {
	detail := map[string]any{
		"code": e.GetErrorCode(),
	}
	if e.GetCause() != nil {
		detail["cause"] = e.GetCause().Error()
	}
	if e.GetErrorDetail() != nil {
		detail["field"] = e.GetErrorDetail().Field
		detail["value"] = e.GetErrorDetail().Value
	}
	return detail
}

func toMapFromError(e error) map[string]any {
	return map[string]any{"cause": e.Error()}
}

func CreateGQLError(ctx context.Context, err error) *gqlerror.Error {
	var cErrs app.CustomErrors
	if errors.As(err, &cErrs) {
		var details []map[string]any
		for _, cErr := range cErrs {
			details = append(details, toMapFromCustomError(cErr))
		}
		return &gqlerror.Error{
			Extensions: map[string]any{"details": details},
		}
	}

	var cErr app.CustomError
	if errors.As(err, &cErr) {
		return &gqlerror.Error{
			Extensions: map[string]any{"details": []map[string]any{toMapFromCustomError(cErr)}},
		}
	}

	re := &gqlerror.Error{
		Extensions: map[string]any{"details": []map[string]any{toMapFromError(err)}},
	}
	return re
}
