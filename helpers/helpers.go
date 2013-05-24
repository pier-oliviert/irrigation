package helpers

import (
	"strconv"
  "time"
)

func Int32ValueFrom(value string, defaultValue int32) int32 {
	parsed, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int32(parsed)
}

func Selected(value bool) string {
	if value == true {
		return "selected"
	}
	return ""
}

func FormatTime(datetime time.Time, format string) string {
  return datetime.Format(format)
}
