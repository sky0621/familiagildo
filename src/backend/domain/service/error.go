package service

import (
	"context"

	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/vo"
)

func Validate[T string | int64](ctx context.Context, valueObjects []vo.ValueObject[T]) app.CustomErrors {
	var customErrors app.CustomErrors
	for _, v := range valueObjects {
		if err := v.Validate(); err != nil {
			customErrors = append(customErrors, app.NewValidationError(ctx, err, v.FieldName(), v.ToVal()))
		}
	}
	return customErrors
}
