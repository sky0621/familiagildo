package gateway

import "errors"

var NoRecords = errors.New("sql: no rows in result set")

func IsNoRecords(err error) bool {
	return err.Error() == NoRecords.Error()
}
