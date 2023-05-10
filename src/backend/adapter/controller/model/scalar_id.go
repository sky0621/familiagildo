package model

import (
	"encoding/base64"
	"fmt"

	"strconv"
	"strings"

	"github.com/sky0621/familiagildo/app"
)

const sep = ":"

func decodeID(id string) (string, int64, error) {
	b, err := base64.RawURLEncoding.DecodeString(id)
	if err != nil {
		return "", 0, app.WrapErrorf(err, "failed to DecodeString[%s]", id)
	}
	items := strings.SplitN(string(b), sep, 2)
	dbUniqueID, err := strconv.ParseInt(items[1], 10, 64)
	if err != nil {
		return items[0], -1, app.WrapErrorf(err, "failed to ParseInt[%s]", items[1])
	}
	return items[0], dbUniqueID, nil
}

func encodeID(typeName string, dbUniqueID int64) string {
	return base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%s%s%d", typeName, sep, dbUniqueID)))
}
