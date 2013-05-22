package helpers

import (
	"strconv"
)

func Int32ValueFrom(value string, defaultValue int32) int32 {
	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int32(parsed)
}
