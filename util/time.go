package util

import (
	"time"
)

var layout = "2006-01-02T15:04:05"

func StringToDate(value string) time.Time {
	convertedTime, _ := time.Parse(layout, value)
	return convertedTime

}
