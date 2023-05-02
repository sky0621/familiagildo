package controller

import (
	"context"
	"github.com/cockroachdb/errors"
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

func CreateGQLError(ctx context.Context, err error) gqlerror.List {
	var cErrs app.CustomErrors
	if errors.As(err, &cErrs) {
		var gErrList gqlerror.List
		for _, cErr := range cErrs {
			gErrList = append(gErrList, &gqlerror.Error{
				Extensions: toMapFromCustomError(cErr),
			})
		}
		return gErrList
	}

	var cErr app.CustomError
	if errors.As(err, &cErr) {
		return gqlerror.List{&gqlerror.Error{
			Extensions: map[string]any{"details": []map[string]any{toMapFromCustomError(cErr)}},
		}}
	}

	re := gqlerror.List{&gqlerror.Error{
		Extensions: map[string]any{"details": []map[string]any{toMapFromError(err)}},
	}}
	return re
}
