package model

import (
	"io"
	"strconv"
	"time"

	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/app/log"

	"github.com/99designs/gqlgen/graphql"
)

const dateLayout = "2006-01-02"

// UnmarshalDate GraphQL -> Domain
func UnmarshalDate(v interface{}) (time.Time, error) {
	s, ok := v.(string)
	if !ok {
		return time.Time{}, app.NewError("not string")
	}
	t, err := time.ParseInLocation(dateLayout, s, app.JST)
	if err != nil {
		return time.Time{}, app.WrapError(err, "failed to ParseInLocation")
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, app.JST), nil
}

// MarshalDate Domain -> GraphQL
func MarshalDate(v time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, err := io.WriteString(w, strconv.Quote(v.Format(dateLayout)))
		if err != nil {
			log.ErrorSend(err)
		}
	})
}
