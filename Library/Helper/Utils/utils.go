package Utils

import (
	"EmployeeService/Constant"
	"context"
	"strings"
	"time"
)

func BeginningOfMonth(date time.Time) time.Time {
	location, _ := time.LoadLocation(Constant.TimeLocation)
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 1, 1, location)
}

func EndOfMonth(date time.Time) time.Time {
	date = date.AddDate(0, 1, -date.Day())
	location, _ := time.LoadLocation(Constant.TimeLocation)
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 1, location)
}

func GetValueOfContext(key string, ctx context.Context) interface{} {
	if ctx.Value(key) != nil {
		return ctx.Value(key)
	}
	return ""
}

func CaseInsensitiveContains(a string, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}
